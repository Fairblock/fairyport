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

type CosmosAccountDetail struct {
	AccNo      uint64
	AccSeqNo   uint64
	AccAddress sdk.AccAddress
	PrivKey    secp256k1.PrivKey
	PubKey     cryptotypes.PubKey
}

func NewCosmosAccount(config config.Config, authClient authTypes.QueryClient) (*CosmosAccountDetail, error) {
	seed, err := bip39.NewSeedWithErrorChecking(config.GetMnemonic(), "")
	if err != nil {
		return nil, err
	}

	master, ch := hd.ComputeMastersFromSeed(seed)
	path := config.CosmosRelayConfig.DerivePath
	priv, err := hd.DerivePrivateKeyForPath(master, ch, path)
	if err != nil {
		return nil, err
	}

	log.Println("Private Key derived successfully: ", hex.EncodeToString(priv))

	privKey := secp256k1.PrivKey{Key: priv}
	pubKey := privKey.PubKey()

	cfg := sdk.GetConfig()
	prefix := config.CosmosRelayConfig.DestinationNode.AccountPrefix
	cfg.SetBech32PrefixForAccount(prefix, prefix+"pub")
	cfg.SetBech32PrefixForValidator(prefix+"valoper", prefix+"valoperpub")
	cfg.SetBech32PrefixForConsensusNode(prefix+"valcons", prefix+"valconspub")

	addr := sdk.AccAddress(pubKey.Address())
	log.Println("Address: ", addr.String())
	accNo, accSeqNo := GetAccountDetails(addr, authClient)
	log.Println("Successfully Fetched Account Details for ", addr)

	return &CosmosAccountDetail{
		PrivKey:    privKey,
		PubKey:     pubKey,
		AccAddress: addr,
		AccNo:      accNo,
		AccSeqNo:   accSeqNo,
	}, nil
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
