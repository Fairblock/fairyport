package grpcservice_test

import (
	"errors"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// MockGRPCDialer allows mocking the grpc.Dial function
type MockGRPCDialer struct {
	mock.Mock
}

func (m *MockGRPCDialer) Dial(endpoint string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	args := m.Called(endpoint, opts)
	return args.Get(0).(*grpc.ClientConn), args.Error(1)
}

func InitializeGRPCServerWithDialer(grpcEndpoint string, dialer func(string, ...grpc.DialOption) (*grpc.ClientConn, error)) *grpc.ClientConn {
	// Create a connection to the gRPC server using the injected dialer.
	conn, err := dialer(
		grpcEndpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

func TestInitializeGRPCServer_Success(t *testing.T) {
	mockDialer := &MockGRPCDialer{}
	expectedConn := &grpc.ClientConn{}

	mockDialer.On("Dial", "localhost:50051", mock.Anything).Return(expectedConn, nil)

	conn := InitializeGRPCServerWithDialer("localhost:50051", mockDialer.Dial)
	assert.NotNil(t, conn)
	assert.Equal(t, expectedConn, conn)

	mockDialer.AssertExpectations(t)
}

func TestInitializeGRPCServer_Failure(t *testing.T) {
	mockDialer := &MockGRPCDialer{}
	mockDialer.On("Dial", "invalid:50051", mock.Anything).Return(nil, errors.New("connection failed"))

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected log.Fatal to terminate the program, but it did not")
		}
	}()

	InitializeGRPCServerWithDialer("invalid:50051", mockDialer.Dial)
	mockDialer.AssertExpectations(t)
}
