package account_test

import (
	"context"
	"testing"

	"github.com/Fairblock/fairyport/config"
	"github.com/Fairblock/fairyport/pkg/cosmos/account"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

// MockAuthQueryClient simulates the authTypes.QueryClient behavior
type MockAuthQueryClient struct {
	mock.Mock
}

func (m *MockAuthQueryClient) Account(ctx context.Context, req *authTypes.QueryAccountRequest, opts ...grpc.CallOption) (*authTypes.QueryAccountResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*authTypes.QueryAccountResponse), args.Error(1)
}

func (m *MockAuthQueryClient) AccountAddressByID(ctx context.Context, req *authTypes.QueryAccountAddressByIDRequest, opts ...grpc.CallOption) (*authTypes.QueryAccountAddressByIDResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*authTypes.QueryAccountAddressByIDResponse), args.Error(1)
}

func (m *MockAuthQueryClient) AccountInfo(ctx context.Context, req *authTypes.QueryAccountInfoRequest, opts ...grpc.CallOption) (*authTypes.QueryAccountInfoResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*authTypes.QueryAccountInfoResponse), args.Error(1)
}

func (m *MockAuthQueryClient) Accounts(ctx context.Context, req *authTypes.QueryAccountsRequest, opts ...grpc.CallOption) (*authTypes.QueryAccountsResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*authTypes.QueryAccountsResponse), args.Error(1)
}

func (m *MockAuthQueryClient) AddressBytesToString(ctx context.Context, req *authTypes.AddressBytesToStringRequest, opts ...grpc.CallOption) (*authTypes.AddressBytesToStringResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*authTypes.AddressBytesToStringResponse), args.Error(1)
}

func (m *MockAuthQueryClient) AddressStringToBytes(ctx context.Context, req *authTypes.AddressStringToBytesRequest, opts ...grpc.CallOption) (*authTypes.AddressStringToBytesResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*authTypes.AddressStringToBytesResponse), args.Error(1)
}

func (m *MockAuthQueryClient) Bech32Prefix(ctx context.Context, req *authTypes.Bech32PrefixRequest, opts ...grpc.CallOption) (*authTypes.Bech32PrefixResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*authTypes.Bech32PrefixResponse), args.Error(1)
}

func (m *MockAuthQueryClient) ModuleAccountByName(ctx context.Context, req *authTypes.QueryModuleAccountByNameRequest, opts ...grpc.CallOption) (*authTypes.QueryModuleAccountByNameResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*authTypes.QueryModuleAccountByNameResponse), args.Error(1)
}

func (m *MockAuthQueryClient) ModuleAccounts(ctx context.Context, req *authTypes.QueryModuleAccountsRequest, opts ...grpc.CallOption) (*authTypes.QueryModuleAccountsResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*authTypes.QueryModuleAccountsResponse), args.Error(1)
}

func (m *MockAuthQueryClient) Params(ctx context.Context, req *authTypes.QueryParamsRequest, opts ...grpc.CallOption) (*authTypes.QueryParamsResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*authTypes.QueryParamsResponse), args.Error(1)
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
