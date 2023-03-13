package app

import (
	"log"

	"github.com/FairBlock/fairyport/config"
	"github.com/FairBlock/fairyport/internal/fairyclient"
	"github.com/FairBlock/fairyport/pkg/account"
	grpcservice "github.com/FairBlock/fairyport/pkg/grpcService"
	cosmosClient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/tx"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
	"google.golang.org/grpc"
)

type App struct {
	Cfg         config.Config
	AccountInfo account.AccountDetails
	GrpcConn    *grpc.ClientConn
	FairyClient *rpchttp.HTTP
	AuthClient  authTypes.QueryClient
	TxClient    tx.ServiceClient
	// Logger      log.Logger
}

func New() *App {
	var cfg config.Config
	cfg.SetConfig()

	conn := grpcservice.InitializeGRPCServer(cfg.GetGRPCEndPoint())
	authClient := authTypes.NewQueryClient(conn)

	// get new client instance from node address
	fairyClient, err := cosmosClient.NewClientFromNode(cfg.GetFairyNodeURI())
	if err != nil {
		log.Fatal(err)
	}

	var accDetails account.AccountDetails
	accDetails.InitializeAccount(cfg, authClient)

	app := &App{
		Cfg:         cfg,
		GrpcConn:    conn,
		AuthClient:  authClient,
		FairyClient: fairyClient,
		AccountInfo: accDetails,
		TxClient:    tx.NewServiceClient(conn),
	}
	return app
}

func (a *App) Start() {
	// start the client
	fairyclient.StartFairyClient(a.FairyClient, &a.AccountInfo, a.TxClient)

	// defer a.FairyClient.Stop()
}
