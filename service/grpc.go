package service

import (
	"log"

	"github.com/FairBlock/fairy-bridge/config"
	"github.com/cosmos/cosmos-sdk/codec"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"google.golang.org/grpc"
)

var grpcConn *grpc.ClientConn
var authClient authTypes.QueryClient

func InitializeGRPCServer(config config.Config) {
	// Create a connection to the gRPC server.
	conn, err := grpc.Dial(
		config.GetGRPCEndPoint(), // your gRPC server address.
		grpc.WithInsecure(),      // The Cosmos SDK doesn't support any transport security mechanism.
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
