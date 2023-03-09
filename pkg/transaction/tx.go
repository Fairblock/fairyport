package transaction

import (
	"context"
	"log"

	"github.com/FairBlock/fairy-bridge/pkg/account"
	fbtypes "github.com/FairBlock/fairy-bridge/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
)

func SendTx(accDetails *account.AccountDetails, txClient tx.ServiceClient, height uint64, data string) error {
	// Choose the codec
	encCfg := simapp.MakeTestEncodingConfig()

	// Create a new TxBuilder.
	txBuilder := encCfg.TxConfig.NewTxBuilder()

	msg := fbtypes.NewMsgCreateAggregatedKeyShare(accDetails.AccAddress.String(), height, data)

	err := txBuilder.SetMsgs(msg)
	if err != nil {
		return err
	}

	txBuilder.SetGasLimit(100000)

	var sigsV2 []signing.SignatureV2
	sigV2 := signing.SignatureV2{
		PubKey: accDetails.PubKey,
		Data: &signing.SingleSignatureData{
			SignMode:  encCfg.TxConfig.SignModeHandler().DefaultMode(),
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
		ChainID:       "destination",
		AccountNumber: accDetails.AccNo,
		Sequence:      accDetails.AccSeqNo,
	}

	sigV2, err = secondSigning(encCfg.TxConfig.SignModeHandler().DefaultMode(), signerData,
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

	if grpcRes.TxResponse.Code == 0 {
		log.Println("successfully Broadcasted Keyshares for height: ", height)
	} else {
		log.Println("Broadcasting Keyshares for Height :", height, " failed with code :", grpcRes.TxResponse.Code)
	}

	// increment sequence number
	accDetails.AccSeqNo = accDetails.AccSeqNo + 1

	return nil
}

func secondSigning(signMode signing.SignMode,
	signerData xauthsigning.SignerData,
	txBuilder client.TxBuilder,
	priv secp256k1.PrivKey,
	txConfig client.TxConfig,
	accSeq uint64) (signing.SignatureV2, error) {
	var sigV2 signing.SignatureV2

	// Generate the bytes to be signed.
	signBytes, err := txConfig.SignModeHandler().GetSignBytes(signMode, signerData, txBuilder.GetTx())
	if err != nil {
		return sigV2, err
	}

	// Sign those bytes
	signature, err := priv.Sign(signBytes)
	if err != nil {
		return sigV2, err
	}

	// Construct the SignatureV2 struct
	sigData := signing.SingleSignatureData{
		SignMode:  signMode,
		Signature: signature,
	}

	sigV2 = signing.SignatureV2{
		PubKey:   priv.PubKey(),
		Data:     &sigData,
		Sequence: accSeq,
	}

	return sigV2, nil
}
