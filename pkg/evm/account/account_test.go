package account_test

import (
	"encoding/hex"
	"testing"

	"github.com/Fairblock/fairyport/pkg/evm/account"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

func TestNewEVMAccount_ValidKey(t *testing.T) {
	// Generate a valid private key for testing
	privateKey, err := crypto.GenerateKey()
	assert.NoError(t, err)

	// Convert the private key to a hex string
	privateKeyBytes := crypto.FromECDSA(privateKey)      // Get the raw private key bytes
	privateKeyHex := hex.EncodeToString(privateKeyBytes) // Encode to a hex string

	// Get the expected address
	expectedAddress := crypto.PubkeyToAddress(privateKey.PublicKey).Hex()

	// Pass the private key hex string to NewEVMAccount
	accountDetail, err := account.NewEVMAccount(privateKeyHex)
	assert.NoError(t, err)
	assert.NotNil(t, accountDetail)
	assert.Equal(t, expectedAddress, accountDetail.Address.Hex())
	assert.NotNil(t, accountDetail.PrivateKey)
}

func TestNewEVMAccount_InvalidHexKey(t *testing.T) {
	// Test with invalid hex string
	invalidKey := "0xINVALIDKEY"

	accountDetail, err := account.NewEVMAccount(invalidKey)
	assert.Error(t, err)
	assert.Nil(t, accountDetail)
}

func TestNewEVMAccount_EmptyKey(t *testing.T) {
	// Test with an empty string
	accountDetail, err := account.NewEVMAccount("")
	assert.Error(t, err)
	assert.Nil(t, accountDetail)
}
