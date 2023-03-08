package service

import (
	"context"
	"log"

	"github.com/FairBlock/fairy-bridge/config"
	"github.com/cosmos/cosmos-sdk/codec"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/gogo/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var grpcConn *grpc.ClientConn
var authClient authTypes.QueryClient

func InitializeGRPCServer(config config.Config) {
	// Create a connection to the gRPC server.
	conn, err := grpc.Dial(
		config.GetGRPCEndPoint(),                                 // gRPC server address.
		grpc.WithTransportCredentials(insecure.NewCredentials()), // The Cosmos SDK doesn't support any transport security mechanism.
		// This instantiates a general gRPC codec which handles proto bytes. We pass in a nil interface registry
		// if the request/response types contain interface instead of 'nil' you should pass the application specific codec.
		grpc.WithDefaultCallOptions(grpc.ForceCodec(codec.NewProtoCodec(nil).GRPCCodec())),
	)
	if err != nil {
		log.Fatal(err)
	}

	grpcConn = conn

	// defer grpcConn.Close()
}

func InitializeAuthClient() {
	// This creates a gRPC client to query the x/bank service.
	authClient = authTypes.NewQueryClient(grpcConn)
}

func GetAccDetails(config config.Config) {
	addr := addr1.String()

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

	// get the account number and sequence number
	accNo = acc.AccountNumber
	seqNo = acc.Sequence
}
