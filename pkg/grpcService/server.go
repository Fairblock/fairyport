package grpcservice

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
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
