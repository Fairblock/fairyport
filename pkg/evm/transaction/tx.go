package transaction

import (
	"context"
	"github.com/Fairblock/fairyport/contract"
	"github.com/Fairblock/fairyport/pkg/evm/account"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

type EVMTxQueue struct {
	ETHClient           ethclient.Client
	ChainID             *big.Int
	FairyringContract   *contract.FairyringContract
	Transactor          *bind.TransactOpts
	TxQueue             chan EVMTx
	UpdateGasEveryBlock bool
}

type EVMTx struct {
	EncryptionKey       []byte
	DecryptionKey       []byte
	DecryptionKeyHeight *big.Int
}

func NewEVMTxQueue(
	rpc,
	contractAddr string,
	accDetail *account.EVMAccountDetail,
	updateGasEveryblock bool,
) (*EVMTxQueue, error) {

	client, err := ethclient.Dial(rpc)
	if err != nil {
		return nil, err
	}

	contractCommonAddr := common.HexToAddress(contractAddr)
	contract, err := contract.NewFairyringContract(contractCommonAddr, client)
	if err != nil {
		return nil, err
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(accDetail.PrivateKey, chainID)
	if err != nil {
		return nil, err
	}

	nonce, err := client.PendingNonceAt(context.Background(), accDetail.Address)
	if err != nil {
		return nil, err
	}

	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.Nonce = big.NewInt(int64(nonce))

	if !updateGasEveryblock {
		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			return nil, err
		}
		auth.GasPrice = gasPrice
	}

	return &EVMTxQueue{
		ETHClient:           *client,
		FairyringContract:   contract,
		ChainID:             chainID,
		Transactor:          auth,
		TxQueue:             make(chan EVMTx, 10),
		UpdateGasEveryBlock: updateGasEveryblock,
	}, nil
}

func (q *EVMTxQueue) QueueEncryptionKey(pubkey []byte, height *big.Int) {
	q.TxQueue <- EVMTx{
		EncryptionKey:       pubkey,
		DecryptionKeyHeight: height,
	}
}

func (q *EVMTxQueue) QueueDecryptionKey(pubkey, aggrKey []byte, aggrKeyHeight *big.Int) {
	q.TxQueue <- EVMTx{
		EncryptionKey:       pubkey,
		DecryptionKey:       aggrKey,
		DecryptionKeyHeight: aggrKeyHeight,
	}
}

func (q *EVMTxQueue) ProcessQueue() {
	var lastSubmittedHeight uint64 = 0
	var lastSubmitEVMHeight uint64 = 0
	for {
		headers := make(chan *types.Header)
		sub, err := q.ETHClient.SubscribeNewHead(context.Background(), headers)
		if err != nil {
			log.Fatalf("Failed Subscribing NewHead event: %v", err)
		}

		log.Printf("Websocket started, subscribed to NewHead Event")

		for {
			select {
			case err := <-sub.Err():
				log.Printf("Error in Subscription: %v", err)
				sub.Unsubscribe()
				break
			case head := <-headers:
				qTx := <-q.TxQueue

				if len(qTx.DecryptionKey) == 0 && len(qTx.EncryptionKey) == 0 {
					continue
				}

				evmHeight := head.Number.Uint64()

				if evmHeight <= lastSubmitEVMHeight {
					log.Println("Already submitted at this EVM height, skip...")
					continue
				} else {
					lastSubmitEVMHeight = evmHeight
				}

				keyHeightUint64 := qTx.DecryptionKeyHeight.Uint64()

				if keyHeightUint64 <= lastSubmittedHeight {
					log.Println("Old key, skip...")
					continue
				} else {
					lastSubmittedHeight = keyHeightUint64
				}

				if q.UpdateGasEveryBlock {
					gasPrice, err := q.ETHClient.SuggestGasPrice(context.Background())
					if err != nil {
						log.Printf("[EVM: %d] [KEY: %d] ERROR getting gas price: %s", evmHeight, keyHeightUint64, err.Error())
						continue
					}

					q.Transactor.GasPrice = gasPrice
				}

				go func(
					ethClient *ethclient.Client,
					tx EVMTx,
					contract *contract.FairyringContract,
					transactor *bind.TransactOpts,
					evmHeight uint64,
				) {
					var (
						sentTx    *types.Transaction
						sendTxErr error
						txName    string
					)

					keyHeight := qTx.DecryptionKeyHeight.Uint64()

					if len(qTx.DecryptionKey) > 0 {
						// Decryption Key not null, submit decryption key
						sentTx, sendTxErr = q.FairyringContract.SubmitDecryptionKey(
							q.Transactor,
							qTx.EncryptionKey,
							qTx.DecryptionKey,
							qTx.DecryptionKeyHeight,
						)
						txName = "SubmitDecryptionKey"
					} else {
						// Decryption key is null, submit encryption key
						sentTx, sendTxErr = q.FairyringContract.SubmitEncryptionKey(
							q.Transactor,
							qTx.EncryptionKey,
						)
						txName = "SubmitEncryptionKey"
					}

					q.Transactor.Nonce = big.NewInt(0).Add(q.Transactor.Nonce, big.NewInt(1))

					if sendTxErr != nil {
						log.Printf("[EVM: %d] [%d] ERROR submitting TX: %s\n", evmHeight, keyHeight, sendTxErr.Error())
						return
					}

					receipt, err := bind.WaitMined(context.Background(), &q.ETHClient, sentTx)
					if err != nil {
						log.Printf("[EVM: %d] [%d] %s FAILED, TXID: %s\n", evmHeight, keyHeight, txName, sentTx.Hash().String())
					} else {
						log.Printf("[EVM: %d] [%d] %s CONFIRMED at [%d], TXID: %s\n", evmHeight, keyHeight, txName, receipt.BlockNumber.Uint64(), receipt.TxHash.String())
					}
				}(&q.ETHClient, qTx, q.FairyringContract, q.Transactor, evmHeight)
			}
		}
	}

}
