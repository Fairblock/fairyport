package account_test

import (
	"context"
	"testing"

	"github.com/Fairblock/fairyport/config"
	"github.com/Fairblock/fairyport/pkg/cosmos/account"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockAuthQueryClient simulates the authTypes.QueryClient behavior
type MockAuthQueryClient struct {
	mock.Mock
}

func (m *MockAuthQueryClient) Account(ctx context.Context, req *authTypes.QueryAccountRequest) (*authTypes.QueryAccountResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*authTypes.QueryAccountResponse), args.Error(1)
}

func TestNewCosmosAccount_Success(t *testing.T) {
	mockAuthClient := &MockAuthQueryClient{}
	mockAuthClient.On("Account", mock.Anything, mock.Anything).Return(&authTypes.QueryAccountResponse{}, nil)

	// Mock configuration
	mockConfig := config.Config{
		CosmosRelayConfig: config.CosmosRelayConfigType{
			DerivePath: "m/44'/118'/0'/0/0",
			Mnemonic:   "convince bike cousin off endless pear prison file where person grace twin multiply teach interest cushion hood vapor twist arrange know april mix artist",
		},
	}

	accountDetail, err := account.NewCosmosAccount(mockConfig, mockAuthClient)
	assert.NoError(t, err)
	assert.NotNil(t, accountDetail)
	assert.Equal(t, uint64(0), accountDetail.AccNo) // Default mocked value
	mockAuthClient.AssertExpectations(t)
}

func TestNewCosmosAccount_InvalidMnemonic(t *testing.T) {
	mockAuthClient := &MockAuthQueryClient{}

	// Mock configuration with an invalid mnemonic
	mockConfig := config.Config{
		CosmosRelayConfig: config.CosmosRelayConfigType{
			DerivePath: "m/44'/118'/0'/0/0",
			Mnemonic:   "invalid-mnemonic",
		},
	}

	accountDetail, err := account.NewCosmosAccount(mockConfig, mockAuthClient)
	assert.Error(t, err)
	assert.Nil(t, accountDetail)
	assert.Contains(t, err.Error(), "invalid mnemonic")
}

func TestNewCosmosAccount_KeyDerivationFailure(t *testing.T) {
	mockAuthClient := &MockAuthQueryClient{}

	// Mock configuration with an invalid derive path
	mockConfig := config.Config{
		CosmosRelayConfig: config.CosmosRelayConfigType{
			DerivePath: "invalid/derive/path",
			Mnemonic:   "convince bike cousin off endless pear prison file where person grace twin multiply teach interest cushion hood vapor twist arrange know april mix artist",
		},
	}

	accountDetail, err := account.NewCosmosAccount(mockConfig, mockAuthClient)
	assert.Error(t, err)
	assert.Nil(t, accountDetail)
	assert.Contains(t, err.Error(), "key derivation failed")
}
