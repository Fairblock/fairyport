package contract_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/Fairblock/fairyport/contract"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockBackend is a mock implementation of the Ethereum backend.
type MockBackend struct {
	mock.Mock
}

func (m *MockBackend) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	args := m.Called(ctx, call, blockNumber)
	return args.Get(0).([]byte), args.Error(1)
}

func (m *MockBackend) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	args := m.Called(ctx, account)
	return args.Get(0).([]byte), args.Error(1)
}

func (m *MockBackend) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	args := m.Called(ctx, account)
	return args.Get(0).(uint64), args.Error(1)
}

func (m *MockBackend) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	args := m.Called(ctx)
	return args.Get(0).(*big.Int), args.Error(1)
}

func (m *MockBackend) EstimateGas(ctx context.Context, call ethereum.CallMsg) (uint64, error) {
	args := m.Called(ctx, call)
	return args.Get(0).(uint64), args.Error(1)
}

func (m *MockBackend) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	args := m.Called(ctx, contract, blockNumber)
	return args.Get(0).([]byte), args.Error(1)
}

func (m *MockBackend) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	args := m.Called(ctx, query)
	return args.Get(0).([]types.Log), args.Error(1)
}

func (m *MockBackend) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	args := m.Called(ctx, number)
	return args.Get(0).(*types.Header), args.Error(1)
}

func (m *MockBackend) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	args := m.Called(ctx, tx)
	return args.Error(0)
}

func (m *MockBackend) SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	args := m.Called(ctx, query, ch)
	return args.Get(0).(ethereum.Subscription), args.Error(1)
}

func (m *MockBackend) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	args := m.Called(ctx)
	return args.Get(0).(*big.Int), args.Error(1)
}

func TestNewFairyringContract(t *testing.T) {
	mockBackend := new(MockBackend)
	address := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678")

	contractInstance, err := contract.NewFairyringContract(address, mockBackend)
	assert.NoError(t, err)
	assert.NotNil(t, contractInstance)
}

func TestFairyringContract_GetRandomnessByHeight(t *testing.T) {
	mockBackend := new(MockBackend)
	address := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678")

	contractInstance, err := contract.NewFairyringContract(address, mockBackend)
	assert.NoError(t, err)

	// Mock backend response
	expectedRandomness := [32]byte{0x1, 0x2, 0x3}
	mockBackend.On("CallContract", mock.Anything, mock.Anything, address).Return(expectedRandomness[:], nil)

	opts := &bind.CallOpts{}
	height := big.NewInt(100)

	randomness, err := contractInstance.FairyringContractCaller.GetRandomnessByHeight(opts, height)
	assert.NoError(t, err)
	assert.Equal(t, expectedRandomness, randomness)
}

func TestFairyringContract_LatestEncryptionKey(t *testing.T) {
	mockBackend := new(MockBackend)
	address := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678")

	contractInstance, err := contract.NewFairyringContract(address, mockBackend)
	assert.NoError(t, err)

	// Mock backend response
	expectedKey := []byte("mockEncryptionKey")
	mockBackend.On("CallContract", mock.Anything, mock.Anything, address).Return(expectedKey, nil)

	opts := &bind.CallOpts{}

	encryptionKey, err := contractInstance.FairyringContractCaller.LatestEncryptionKey(opts)
	assert.NoError(t, err)
	assert.Equal(t, expectedKey, encryptionKey)
}

func TestFairyringContract_LatestRandomness(t *testing.T) {
	mockBackend := new(MockBackend)
	address := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678")

	contractInstance, err := contract.NewFairyringContract(address, mockBackend)
	assert.NoError(t, err)

	// Mock backend response
	expectedRandomness := [32]byte{0x1, 0x2, 0x3}
	expectedHeight := big.NewInt(200)
	mockBackend.On("CallContract", mock.Anything, mock.Anything, address).Return(append(expectedRandomness[:], expectedHeight.Bytes()...), nil)

	opts := &bind.CallOpts{}

	randomness, height, err := contractInstance.FairyringContractCaller.LatestRandomness(opts)
	assert.NoError(t, err)
	assert.Equal(t, expectedRandomness, randomness)
	assert.Equal(t, expectedHeight, height)
}
