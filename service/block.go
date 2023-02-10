package service

import (
	cosmosClient "github.com/cosmos/cosmos-sdk/client"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	tendermintTypes "github.com/tendermint/tendermint/types"
)

func processBlock(block tendermintTypes.EventDataNewBlock) {
	//fairyHeight := block.Block.Header.Height

}

func sendTx() error {
	// Choose the codec
	encCfg := simapp.MakeTestEncodingConfig()
	kb := cosmosClient.Context.Keyring

	// Create a new TxBuilder.
	txBuilder := encCfg.TxConfig.NewTxBuilder()

	priv1, _, addr1 := testdata.KeyTestPubAddr()

	return nil
}

func getKeysFromMnemonic() (cryptotypes.PrivKey, cryptotypes.PubKey, sdk.AccAddress) {

}
