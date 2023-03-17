package grpcservice

import (
	"log"

	//	"github.com/cosmos/cosmos-sdk/codec"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitializeGRPCServer(grpcEndpoint string) *grpc.ClientConn {
	// Create a connection to the gRPC server.
	conn, err := grpc.Dial(
		grpcEndpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}

	return conn
}
