package fairyport_app_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Fairblock/fairyport/config"
	"github.com/Fairblock/fairyport/internal/fairyport_app"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	ctypes "github.com/cometbft/cometbft/rpc/core/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock dependencies
type MockFairyClient struct {
	*rpchttp.HTTP
	mock.Mock
}

func (m *MockFairyClient) Start() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockFairyClient) Stop() {}

func (m *MockFairyClient) Subscribe(ctx context.Context, subscriber, query string) (<-chan ctypes.ResultEvent, error) {
	args := m.Called(ctx, subscriber, query)
	return args.Get(0).(<-chan ctypes.ResultEvent), args.Error(1)
}

func TestNewFairyportApp_Success(t *testing.T) {
	mockCfg := config.Config{
		CosmosRelayConfig: config.CosmosRelayConfigType{
			MetricsPort: 9090,
			DestinationNode: config.Node{
				ChainId: "cosmoshub-4",
			},
		},
	}

	app := fairyport_app.NewFairyportApp(mockCfg)
	assert.NotNil(t, app)
	assert.NotNil(t, app.GrpcConn)
	assert.NotNil(t, app.AuthClient)
	assert.NotNil(t, app.FairyClient)
}

func TestNewFairyportApp_EVMAccountError(t *testing.T) {
	mockCfg := config.Config{}

	// Ensure EVM account error handling doesn't crash the application
	app := fairyport_app.NewFairyportApp(mockCfg)
	assert.NotNil(t, app)
	assert.Nil(t, app.EVMSenderAddress)
	assert.False(t, app.RelayToEVMs)
}

func TestStartFairyport_Success(t *testing.T) {
	mockFairyClient := &MockFairyClient{}
	mockCfg := config.Config{
		CosmosRelayConfig: config.CosmosRelayConfigType{
			MetricsPort: 9090,
		},
	}
	app := &fairyport_app.FairyportApp{
		Cfg:         mockCfg,
		FairyClient: mockFairyClient.HTTP,
	}

	// Mock client behavior
	mockFairyClient.On("Start").Return(nil)
	mockFairyClient.On("Subscribe", mock.Anything, "fairyport", "tm.event = 'Tx'").Return(make(chan ctypes.ResultEvent), nil)

	go func() {
		time.Sleep(100 * time.Millisecond) // Allow some execution time
		app.FairyClient.Stop()
	}()

	assert.NotPanics(t, func() {
		app.StartFairyport()
	})
	mockFairyClient.AssertExpectations(t)
}

func TestStartFairyport_SubscribeError(t *testing.T) {
	mockFairyClient := &MockFairyClient{}
	mockCfg := config.Config{}
	app := &fairyport_app.FairyportApp{
		Cfg:         mockCfg,
		FairyClient: mockFairyClient.HTTP,
	}

	// Mock client behavior
	mockFairyClient.On("Start").Return(nil)
	mockFairyClient.On("Subscribe", mock.Anything, "fairyport", "tm.event = 'Tx'").Return(nil, errors.New("subscription error"))

	assert.Panics(t, func() {
		app.StartFairyport()
	})
	mockFairyClient.AssertExpectations(t)
}
