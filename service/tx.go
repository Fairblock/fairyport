package service

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"

	fbtypes "github.com/FairBlock/fairy-bridge/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"

	"github.com/FairBlock/fairy-bridge/config"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/cosmos/go-bip39"
)

var priv1 secp256k1.PrivKey
var addr1 sdk.AccAddress
var seqNo uint64 = 0

func InitializeAccount(config config.Config) {
	seed, err := bip39.NewSeedWithErrorChecking(config.GetMnemonic(), "")
	if err != nil {
		log.Fatal(err)
	}

	master, ch := hd.ComputeMastersFromSeed(seed)
	path := "m/44'/118'/0'/0/0"
	priv, err := hd.DerivePrivateKeyForPath(master, ch, path)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Derivation Path: ", path)
	fmt.Println("Private Key: ", hex.EncodeToString(priv))

	// cryptotypes.PrivKey(priv)
	priv1 = secp256k1.PrivKey{Key: priv}
	pub := priv1.PubKey()
	addr1 = sdk.AccAddress(pub.Address())

	// priv1, _, addr1 = testdata.KeyTestPubAddr()
	fmt.Println(addr1.String())
}

func sendTx(height uint64, data string) error {
	// Choose the codec
	encCfg := simapp.MakeTestEncodingConfig()

	// Create a new TxBuilder.
	txBuilder := encCfg.TxConfig.NewTxBuilder()

	msg := fbtypes.NewMsgCreateAggregatedKeyShare(addr1.String(), height, data)

	err := txBuilder.SetMsgs(msg)
	if err != nil {
		fmt.Println("1: ", err)
		return err
	}

	txBuilder.SetGasLimit(100000)
	// txBuilder.SetFeeAmount(...)
	// txBuilder.SetMemo(...)
	// txBuilder.SetTimeoutHeight(...)

	var sigsV2 []signing.SignatureV2
	sigV2 := signing.SignatureV2{
		PubKey: priv1.PubKey(),
		Data: &signing.SingleSignatureData{
			SignMode:  encCfg.TxConfig.SignModeHandler().DefaultMode(),
			Signature: nil,
		},
		Sequence: seqNo,
	}

	sigsV2 = append(sigsV2, sigV2)

	err = txBuilder.SetSignatures(sigsV2...)
	if err != nil {
		fmt.Println("2: ", err)
		return err
	}

	sigsV2 = []signing.SignatureV2{}
	signerData := xauthsigning.SignerData{
		ChainID:       "destination",
		AccountNumber: 1,
		Sequence:      seqNo,
	}

	sigV2, err = secondSigning(encCfg.TxConfig.SignModeHandler().DefaultMode(), signerData,
		txBuilder, priv1, encCfg.TxConfig, seqNo)

	if err != nil {
		fmt.Println("2.1: ", err)
		return err
	}

	sigsV2 = append(sigsV2, sigV2)

	err = txBuilder.SetSignatures(sigsV2...)
	if err != nil {
		fmt.Println("2.2: ", err)
		return err
	}

	// Generated Protobuf-encoded bytes.
	txBytes, err := encCfg.TxConfig.TxEncoder()(txBuilder.GetTx())
	if err != nil {
		fmt.Println("3: ", err)
		return err
	}

	fmt.Println(len(txBytes))

	txClient := tx.NewServiceClient(grpcConn)
	// We then call the BroadcastTx method on this client.
	grpcRes, err := txClient.BroadcastTx(
		context.Background(),
		&tx.BroadcastTxRequest{
			Mode:    tx.BroadcastMode_BROADCAST_MODE_SYNC,
			TxBytes: txBytes, // Proto-binary of the signed transaction
		},
	)
	if err != nil {
		fmt.Println("5: ", err)
		return err
	}

	fmt.Println(grpcRes.TxResponse)
	// fmt.Println(grpcRes.TxResponse.Code) // Should be `0` if the tx is successful
	seqNo = seqNo + 1

	return nil
}

// func getAccSeq() {
// 	rsp, err := authClient.Account(
// 		context.Background(),
// 		&types.QueryAccountRequest{
// 			Address: addr1.String(),
// 		},
// 	)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var acc types.AccountI
// 	acc = rsp.Account.(types.AccountI)
// }

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
