package service

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"

	fbtypes "github.com/FairBlock/fairy-bridge/types"
	"github.com/tendermint/tendermint/crypto/secp256k1"

	"github.com/FairBlock/fairy-bridge/config"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/go-bip39"
	tendermintTypes "github.com/tendermint/tendermint/types"
)

var priv1 secp256k1.PrivKey
var addr1 sdk.AccAddress

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

	priv1 = secp256k1.PrivKey(priv)
	pub := priv1.PubKey()
	addr1 = sdk.AccAddress(pub.Address())

	// priv1, _, addr1 = testdata.KeyTestPubAddr()
	fmt.Println(addr1.String())
}

func processBlock(block tendermintTypes.EventDataNewBlock) {
	//fairyHeight := block.Block.Header.Height

}

func sendTx(height string) error {
	// Choose the codec
	encCfg := simapp.MakeTestEncodingConfig()

	// Create a new TxBuilder.
	txBuilder := encCfg.TxConfig.NewTxBuilder()

	msg := fbtypes.NewMsgRegisterHeight(addr1.String(), height)

	err := txBuilder.SetMsgs(msg)
	if err != nil {
		return err
	}

	// txBuilder.SetGasLimit(...)
	// txBuilder.SetFeeAmount(...)
	// txBuilder.SetMemo(...)
	// txBuilder.SetTimeoutHeight(...)

	rsp, err := authClient.Account(
		context.Background(),
		&types.QueryAccountRequest{
			Address: addr1.String(),
		},
	)

	rsp.

	return nil
}

func getKeysFromMnemonic() (cryptotypes.PrivKey, cryptotypes.PubKey, sdk.AccAddress) {
	return testdata.KeyTestPubAddr()
}
