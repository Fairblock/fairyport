package transaction

import (
	"context"
	"errors"
	"github.com/Fairblock/fairyport/pkg/cosmos/account"
	fbtypes "github.com/Fairblock/fairyring/x/pep/types"
	"github.com/cosmos/cosmos-sdk/client"
	clienttx "github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/skip-mev/block-sdk/v2/testutils"
	"strings"
	"time"
)

func SendTx(
	accDetails *account.CosmosAccountDetail,
	txClient tx.ServiceClient,
	height uint64,
	data, destinationChainID string) error {
	// Choose the codec
	encCfg := testutils.CreateTestEncodingConfig()

	// Create a new TxBuilder.
	txBuilder := encCfg.TxConfig.NewTxBuilder()

	msg := fbtypes.NewMsgSubmitDecryptionKey(accDetails.AccAddress.String(), height, data)

	err := txBuilder.SetMsgs(msg)
	if err != nil {
		return err
	}

	txBuilder.SetGasLimit(100000)

	var sigsV2 []signing.SignatureV2
	sigV2 := signing.SignatureV2{
		PubKey: accDetails.PubKey,
		Data: &signing.SingleSignatureData{
			SignMode:  1,
			Signature: nil,
		},
		Sequence: accDetails.AccSeqNo,
	}

	sigsV2 = append(sigsV2, sigV2)

	err = txBuilder.SetSignatures(sigsV2...)
	if err != nil {
		return err
	}

	sigsV2 = []signing.SignatureV2{}
	signerData := xauthsigning.SignerData{
		ChainID:       destinationChainID,
		AccountNumber: accDetails.AccNo,
		Sequence:      accDetails.AccSeqNo,
		PubKey:        accDetails.PubKey,
		Address:       accDetails.AccAddress.String(),
	}

	sigV2, err = secondSigning(signerData,
		txBuilder, accDetails.PrivKey, encCfg.TxConfig, accDetails.AccSeqNo)

	if err != nil {
		return err
	}

	sigsV2 = append(sigsV2, sigV2)

	err = txBuilder.SetSignatures(sigsV2...)
	if err != nil {
		return err
	}

	// Generated Protobuf-encoded bytes.
	txBytes, err := encCfg.TxConfig.TxEncoder()(txBuilder.GetTx())
	if err != nil {
		return err
	}

	// We then call the BroadcastTx method on this client.
	grpcRes, err := txClient.BroadcastTx(
		context.Background(),
		&tx.BroadcastTxRequest{
			Mode:    tx.BroadcastMode_BROADCAST_MODE_SYNC,
			TxBytes: txBytes, // Proto-binary of the signed transaction
		},
	)
	if err != nil {
		return err
	}

	txResp, err := WaitForTx(txClient, grpcRes.TxResponse.TxHash, time.Second)
	if err != nil {
		return err
	}

	// increment sequence number
	accDetails.AccSeqNo = accDetails.AccSeqNo + 1

	if txResp.TxResponse.Code != 0 {
		return errors.New(txResp.TxResponse.RawLog)
	}

	return nil
}

func secondSigning(
	signerData xauthsigning.SignerData,
	txBuilder client.TxBuilder,
	priv secp256k1.PrivKey,
	txConfig client.TxConfig,
	accSeq uint64) (signing.SignatureV2, error) {
	var sigV2 signing.SignatureV2

	sigV2, err := clienttx.SignWithPrivKey(
		context.Background(), 1, signerData, txBuilder, &priv,
		txConfig, accSeq,
	)

	return sigV2, err
}

func WaitForTx(txClient tx.ServiceClient, hash string, rate time.Duration) (*tx.GetTxResponse, error) {
	for {
		resp, err := txClient.GetTx(context.Background(), &tx.GetTxRequest{Hash: hash})
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				time.Sleep(rate)
				continue
			}

			return nil, err
		}
		return resp, nil
	}
}
