package fairyport_app_test

import (
	"errors"
	"testing"

	"github.com/Fairblock/fairyport/config"
	"github.com/Fairblock/fairyport/internal/fairyport_app"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockTransaction simulates the behavior of the SendTx function
type MockTransaction struct {
	mock.Mock
}

func (m *MockTransaction) SendTx(accInfo interface{}, txClient interface{}, height uint64, key, chainID string) error {
	args := m.Called(accInfo, txClient, height, key, chainID)
	return args.Error(0)
}

func TestRelayDecryptionKeyToCosmos_Success(t *testing.T) {
	// Mock app and dependencies
	mockTx := &MockTransaction{}
	mockApp := &fairyport_app.FairyportApp{
		CosmosAccountInfo: nil, // Replace with appropriate mock or data
		TxClient:          nil, // Replace with appropriate mock or data
		Cfg: config.Config{
			CosmosRelayConfig: config.CosmosRelayConfigType{
				DestinationNode: config.Node{
					ChainId: "cosmoshub-4",
				},
			},
		},
	}

	mockTx.On("SendTx", mock.Anything, mock.Anything, uint64(100), "mockKey", "cosmoshub-4").Return(nil)

	// Inject mock
	// transaction.SendTx = mockTx.SendTx

	err := fairyport_app.RelayDecryptionKeyToCosmos(mockApp, 100, "mockKey")
	assert.NoError(t, err)
	mockTx.AssertExpectations(t)
}

func TestRelayDecryptionKeyToCosmos_Failure(t *testing.T) {
	// Mock app and dependencies
	mockTx := &MockTransaction{}
	mockApp := &fairyport_app.FairyportApp{
		CosmosAccountInfo: nil, // Replace with appropriate mock or data
		TxClient:          nil, // Replace with appropriate mock or data
		Cfg: config.Config{
			CosmosRelayConfig: config.CosmosRelayConfigType{
				DestinationNode: config.Node{
					ChainId: "cosmoshub-4",
				},
			},
		},
	}

	mockTx.On("SendTx", mock.Anything, mock.Anything, uint64(100), "mockKey", "cosmoshub-4").Return(errors.New("transaction failed"))

	// Inject mock
	// transaction.SendTx = mockTx.SendTx

	err := fairyport_app.RelayDecryptionKeyToCosmos(mockApp, 100, "mockKey")
	assert.Error(t, err)
	assert.Equal(t, "transaction failed", err.Error())
	mockTx.AssertExpectations(t)
}
