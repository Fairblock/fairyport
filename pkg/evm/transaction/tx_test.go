package transaction_test

import (
	"context"
	"errors"
	"math/big"
	"testing"

	"github.com/Fairblock/fairyport/contract"
	"github.com/Fairblock/fairyport/pkg/evm/account"
	"github.com/Fairblock/fairyport/pkg/evm/transaction"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockClient simulates the behavior of ethclient.Client
type MockClient struct {
	mock.Mock
}

func (m *MockClient) Dial(url string) (*ethclient.Client, error) {
	args := m.Called(url)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*ethclient.Client), args.Error(1)
}

func (m *MockClient) ChainID(ctx context.Context) (*big.Int, error) {
	args := m.Called(ctx)
	return args.Get(0).(*big.Int), args.Error(1)
}

func (m *MockClient) PendingNonceAt(ctx context.Context, addr common.Address) (uint64, error) {
	args := m.Called(ctx, addr)
	return args.Get(0).(uint64), args.Error(1)
}

func (m *MockClient) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	args := m.Called(ctx)
	return args.Get(0).(*big.Int), args.Error(1)
}

// MockFairyringContract simulates contract behavior
type MockFairyringContract struct {
	mock.Mock
}

func (m *MockFairyringContract) NewFairyringContract(addr common.Address, client *ethclient.Client) (*contract.FairyringContract, error) {
	args := m.Called(addr, client)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*contract.FairyringContract), args.Error(1)
}

func TestNewEVMTxQueue_Success(t *testing.T) {
	mockClient := &MockClient{}
	mockFairyringContract := &MockFairyringContract{}

	rpc := "http://mock-rpc"
	contractAddr := "0xMockContractAddress"
	accDetail := &account.EVMAccountDetail{
		PrivateKey: mockPrivateKey(), // Replace with valid private key for test
		Address:    common.HexToAddress("0xMockAddress"),
	}
	updateGasEveryBlock := false

	mockClient.On("Dial", rpc).Return(&ethclient.Client{}, nil)
	mockFairyringContract.On("NewFairyringContract", common.HexToAddress(contractAddr), mockClient).Return(&contract.FairyringContract{}, nil)
	mockClient.On("ChainID", mock.Anything).Return(big.NewInt(1), nil)
	mockClient.On("PendingNonceAt", mock.Anything, accDetail.Address).Return(uint64(0), nil)
	mockClient.On("SuggestGasPrice", mock.Anything).Return(big.NewInt(1000000000), nil)

	queue, err := transaction.NewEVMTxQueue(rpc, contractAddr, accDetail, updateGasEveryBlock)

	assert.NoError(t, err)
	assert.NotNil(t, queue)
	assert.Equal(t, queue.UpdateGasEveryBlock, updateGasEveryBlock)
	assert.Equal(t, queue.ChainID.Cmp(big.NewInt(1)), 0)
	mockClient.AssertExpectations(t)
	mockFairyringContract.AssertExpectations(t)
}

func TestNewEVMTxQueue_DialFailure(t *testing.T) {
	mockClient := &MockClient{}

	rpc := "http://mock-rpc"
	contractAddr := "0xMockContractAddress"
	accDetail := &account.EVMAccountDetail{
		PrivateKey: mockPrivateKey(), // Replace with valid private key for test
		Address:    common.HexToAddress("0xMockAddress"),
	}
	updateGasEveryBlock := false

	mockClient.On("Dial", rpc).Return(nil, errors.New("failed to connect"))

	_, err := transaction.NewEVMTxQueue(rpc, contractAddr, accDetail, updateGasEveryBlock)

	assert.Error(t, err)
	assert.Equal(t, err.Error(), "failed to connect")
	mockClient.AssertExpectations(t)
}
