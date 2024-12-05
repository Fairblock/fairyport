package fairyport_app

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"time"

	"github.com/Fairblock/fairyport/config"
	"github.com/Fairblock/fairyport/contract"
	"github.com/Fairblock/fairyport/internal/events"
	cosmosAccount "github.com/Fairblock/fairyport/pkg/cosmos/account"
	grpcservice "github.com/Fairblock/fairyport/pkg/cosmos/grpcService"
	evmAccount "github.com/Fairblock/fairyport/pkg/evm/account"
	evmTx "github.com/Fairblock/fairyport/pkg/evm/transaction"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	cosmosClient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/tx"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

var (
	failedBroadcastAggregatedKeyShare = promauto.NewCounter(prometheus.CounterOpts{
		Name: "fairyport_failed_broadcast_aggregated_keyshare",
		Help: "The total number of failed key share generated",
	})

	validBroadcastAggregatedKeyShare = promauto.NewCounter(prometheus.CounterOpts{
		Name: "fairyport_valid_broadcast_aggregated_keyshare",
		Help: "The total number of valid key share generated",
	})
)

type FairyportApp struct {
	Cfg               config.Config
	CosmosAccountInfo *cosmosAccount.CosmosAccountDetail
	EVMSenderAddress  *common.Address
	GrpcConn          *grpc.ClientConn
	FairyClient       *rpchttp.HTTP
	AuthClient        authTypes.QueryClient
	TxClient          tx.ServiceClient
	EVMTxQueue        *evmTx.EVMTxQueue
	RelayToEVMs       bool
	RelayToCosmos     bool
}

type EVMRelayContractTarget struct {
	ETHClient         *ethclient.Client
	ChainID           *big.Int
	FairyringContract *contract.FairyringContract
	Transactor        *bind.TransactOpts
}

func NewFairyportApp(cfg config.Config) *FairyportApp {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	conn := grpcservice.InitializeGRPCServer(cfg.GetGRPCEndPoint())
	authClient := authTypes.NewQueryClient(conn)

	// get new client instance from node address
	fairyClient, err := cosmosClient.NewClientFromNode(cfg.GetFairyNodeURI())
	if err != nil {
		log.Fatalf("Error creating connection to Fairyring node, error: %s", err.Error())
	}

	accDetails, cosmosErr := cosmosAccount.NewCosmosAccount(os.Getenv("COSMOS_MNEMONIC"), cfg, authClient)
	if cosmosErr != nil {
		log.Printf("Unable to initialize Cosmos Account, Disabled relaying keys to destination cosmos chain, err: %s", cosmosErr.Error())
	} else {
		log.Printf("Cosmos Account Initialized, Address: %s.", accDetails.AccAddress.String())
	}

	var address *common.Address
	var evmTxQueue *evmTx.EVMTxQueue

	evmAcc, evmErr := evmAccount.NewEVMAccount(os.Getenv("EVM_PKEY"))

	if evmErr != nil {
		log.Printf("Unable to initialize EVM Account, Disabled relaying keys to destination evm chains, err: %s", evmErr.Error())
		address = nil
		evmTxQueue = nil
	} else {
		address = &evmAcc.Address
		log.Printf("EVM Account Initialized, Address: %s.", address.String())
		q, err := evmTx.NewEVMTxQueue(
			cfg.EVMRelayTarget.ChainRPC,
			cfg.EVMRelayTarget.ContractAddress,
			evmAcc,
			false,
		)
		if err != nil {
			log.Printf("Unable to initialize EVM Tx Queue, Relay keys to EVM chain disabled. Error: %s", err.Error())
		}
		evmTxQueue = q
	}

	return &FairyportApp{
		Cfg:               cfg,
		GrpcConn:          conn,
		AuthClient:        authClient,
		FairyClient:       fairyClient,
		CosmosAccountInfo: accDetails,
		EVMSenderAddress:  address,
		RelayToEVMs:       evmErr == nil && evmTxQueue != nil,
		RelayToCosmos:     cosmosErr == nil,
		TxClient:          tx.NewServiceClient(conn),
		EVMTxQueue:        evmTxQueue,
	}
}

func (app *FairyportApp) StartFairyport() {
	if app == nil {
		log.Fatal("App is nil")
	}
	err := app.FairyClient.Start()
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer app.FairyClient.Stop()
	defer cancel()

	// subscribe to new blocks
	rsp, err := app.FairyClient.Subscribe(ctx, "fairyport", "tm.event = 'Tx'")
	if err != nil {
		log.Fatalf("Unable to subscribe to new block event, error: %s", err.Error())
	}

	http.Handle("/metrics", promhttp.Handler())
	log.Printf("Metrics Started, Listening at port: %d\n", app.Cfg.CosmosRelayConfig.MetricsPort)
	go http.ListenAndServe(fmt.Sprintf(":%d", app.Cfg.CosmosRelayConfig.MetricsPort), nil)
	if app.EVMTxQueue != nil {
		go app.EVMTxQueue.ProcessQueue()
		log.Println("EVM Tx Queue Handler Started")
	}

	sentPubkey := make(map[string]bool)

	for data := range rsp {
		// process the events
		height, aggregatedKeyShare, pubkey, err := events.ProcessEvents(data.Events)
		if err != nil {
			continue
		}

		if app.RelayToCosmos {
			err := RelayDecryptionKeyToCosmos(app, height, aggregatedKeyShare)
			if err != nil {
				failedBroadcastAggregatedKeyShare.Inc()
				log.Printf("[Cosmos] [%d] | Failed to submit decryption key. Error: %s", height, err.Error())
				continue
			}
			validBroadcastAggregatedKeyShare.Inc()
			log.Printf("[Cosmos] [%d] Successfully Broadcasted Decryption key <%s> to Cosmos Chain", height, aggregatedKeyShare)
		}

		if app.RelayToEVMs {

			decryptionKeyByte, err := hex.DecodeString(aggregatedKeyShare)
			if err != nil {
				log.Printf("[EVM] [%d] Error decoding decryption key to bytes, error: %s", height, err.Error())
				continue
			}

			decryptionKeyBigHeight := big.NewInt(int64(height))

			pubKeyBytes, err := hex.DecodeString(pubkey)
			if err != nil {
				log.Printf("[EVM] [%d] Error decoding public key to bytes, error: %s", height, err.Error())
				continue
			}

			if !sentPubkey[pubkey] {
				sentPubkey[pubkey] = true
				app.EVMTxQueue.QueueEncryptionKey(pubKeyBytes, decryptionKeyBigHeight)
			}

			app.EVMTxQueue.QueueDecryptionKey(pubKeyBytes, decryptionKeyByte, decryptionKeyBigHeight)
		}
	}
}
