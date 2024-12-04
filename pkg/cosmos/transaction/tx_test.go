package transaction_test

import (
	"context"
	"errors"
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/Fairblock/fairyport/pkg/cosmos/account"
	"github.com/Fairblock/fairyport/pkg/cosmos/transaction"
	"google.golang.org/grpc"
)

type MockTxServiceClient struct {
	mock.Mock
}

func (m *MockTxServiceClient) BroadcastTx(ctx context.Context, req *tx.BroadcastTxRequest, opts ...grpc.CallOption) (*tx.BroadcastTxResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*tx.BroadcastTxResponse), args.Error(1)
}

func (m *MockTxServiceClient) GetTx(ctx context.Context, req *tx.GetTxRequest, opts ...grpc.CallOption) (*tx.GetTxResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*tx.GetTxResponse), args.Error(1)
}

func (m *MockTxServiceClient) GetTxsEvent(ctx context.Context, req *tx.GetTxsEventRequest, opts ...grpc.CallOption) (*tx.GetTxsEventResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*tx.GetTxsEventResponse), args.Error(1)
}

func (m *MockTxServiceClient) GetBlockWithTxs(ctx context.Context, req *tx.GetBlockWithTxsRequest, opts ...grpc.CallOption) (*tx.GetBlockWithTxsResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*tx.GetBlockWithTxsResponse), args.Error(1)
}

func (m *MockTxServiceClient) Simulate(ctx context.Context, req *tx.SimulateRequest, opts ...grpc.CallOption) (*tx.SimulateResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*tx.SimulateResponse), args.Error(1)
}

func (m *MockTxServiceClient) TxDecode(ctx context.Context, req *tx.TxDecodeRequest, opts ...grpc.CallOption) (*tx.TxDecodeResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*tx.TxDecodeResponse), args.Error(1)
}

func (m *MockTxServiceClient) TxDecodeAmino(ctx context.Context, req *tx.TxDecodeAminoRequest, opts ...grpc.CallOption) (*tx.TxDecodeAminoResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*tx.TxDecodeAminoResponse), args.Error(1)
}

func (m *MockTxServiceClient) TxEncode(ctx context.Context, req *tx.TxEncodeRequest, opts ...grpc.CallOption) (*tx.TxEncodeResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*tx.TxEncodeResponse), args.Error(1)
}

func (m *MockTxServiceClient) TxEncodeAmino(ctx context.Context, req *tx.TxEncodeAminoRequest, opts ...grpc.CallOption) (*tx.TxEncodeAminoResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*tx.TxEncodeAminoResponse), args.Error(1)
}

func TestSendTx_Success(t *testing.T) {
	// Mock account details
	accAddress, err := sdk.AccAddressFromBech32("cosmos1exampleaddress")
	assert.NoError(t, err)

	accDetails := &account.CosmosAccountDetail{
		AccAddress: accAddress,
		PrivKey:    secp256k1.PrivKey{Key: []byte("mockprivatekey")},
	}

	// Mock client and transaction response
	mockTxClient := &MockTxServiceClient{}
	mockTxClient.On("BroadcastTx", mock.Anything, mock.Anything).Return(
		&tx.BroadcastTxResponse{
			TxResponse: &sdk.TxResponse{Code: 0, RawLog: "Success"},
		},
		nil,
	)

	err = transaction.SendTx(accDetails, mockTxClient, 100, "mockData", "destinationChainID")
	assert.NoError(t, err)

	mockTxClient.AssertExpectations(t)
}

func TestSendTx_BroadcastTxError(t *testing.T) {
	// Mock account details
	accAddress, err := sdk.AccAddressFromBech32("cosmos1exampleaddress")
	assert.NoError(t, err)

	accDetails := &account.CosmosAccountDetail{
		AccAddress: accAddress,
		PrivKey:    secp256k1.PrivKey{Key: []byte("mockprivatekey")},
	}

	// Mock client with an error response
	mockTxClient := &MockTxServiceClient{}
	mockTxClient.On("BroadcastTx", mock.Anything, mock.Anything).Return(
		nil,
		errors.New("broadcast failed"),
	)

	err = transaction.SendTx(accDetails, mockTxClient, 100, "mockData", "destinationChainID")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "broadcast failed")

	mockTxClient.AssertExpectations(t)
}
