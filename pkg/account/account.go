package account

import (
	"context"
	"encoding/hex"
	"log"

	"github.com/Fairblock/fairyport/config"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/go-bip39"
	"github.com/gogo/protobuf/proto"
)

type AccountDetails struct {
	AccNo      uint64
	AccSeqNo   uint64
	AccAddress sdk.AccAddress
	PrivKey    secp256k1.PrivKey
	PubKey     cryptotypes.PubKey
}

func (a *AccountDetails) InitializeAccount(config config.Config, authClient authTypes.QueryClient) {
	seed, err := bip39.NewSeedWithErrorChecking(config.GetMnemonic(), "")
	if err != nil {
		log.Fatal(err)
	}

	master, ch := hd.ComputeMastersFromSeed(seed)
	path := config.DerivePath
	priv, err := hd.DerivePrivateKeyForPath(master, ch, path)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Private Key derived successfully: ", hex.EncodeToString(priv))

	a.PrivKey = secp256k1.PrivKey{Key: priv}
	a.PubKey = a.PrivKey.PubKey()

	cfg := sdk.GetConfig()
	prefix := config.DestinationNode.AccountPrefix
	cfg.SetBech32PrefixForAccount(prefix, prefix+"pub")
	cfg.SetBech32PrefixForValidator(prefix+"valoper", prefix+"valoperpub")
	cfg.SetBech32PrefixForConsensusNode(prefix+"valcons", prefix+"valconspub")

	a.AccAddress = sdk.AccAddress(a.PubKey.Address())
	log.Println("Address: ", a.AccAddress.String())
	a.AccNo, a.AccSeqNo = GetAccountDetails(a.AccAddress, authClient)
	log.Println("Successfully Fetched Account Details for ", a.AccAddress)
}

func GetAccountDetails(address sdk.AccAddress, authClient authTypes.QueryClient) (uint64, uint64) {
	addr := address.String()

	// create a QueryAccountRequest to send to the server
	req := &authTypes.QueryAccountRequest{
		Address: addr,
	}

	// send the request to the server
	resp, err := authClient.Account(context.Background(), req)
	if err != nil {
		log.Fatalf("failed to get account: %v", err)
	}

	// decode the account data from the response
	var acc authTypes.BaseAccount
	err = proto.Unmarshal(resp.Account.Value, &acc)
	if err != nil {
		log.Fatalf("failed to decode account: %v", err)
	}

	return acc.AccountNumber, acc.Sequence
}
