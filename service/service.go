package service

import (
	"context"
	"fmt"
	"time"

	"github.com/FairBlock/fairy-bridge/config"
	cosmosClient "github.com/cosmos/cosmos-sdk/client"
)

func NewSerice() error {
	// Set configuration from config.yml file
	var config config.Config
	config.SetConfig()

	// Initialize account
	InitializeAccount(config)

	// Initialize GRPC Server
	InitializeGRPCServer(config)

	fairyNodeURI := config.GetFairyNodeURI()

	// get new client instance from node address
	fairyClient, err := cosmosClient.NewClientFromNode(fairyNodeURI)
	if err != nil {
		panic(err)
	}

	// start the client
	err = fairyClient.Start()
	if err != nil {
		panic(err)
	}
	fmt.Println("started client: listening to ", fairyNodeURI)

	defer fairyClient.Stop()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// query for new block event
	query := "tm.event = 'NewBlock'"

	// subscribe to new blocks
	rsp, err := fairyClient.Subscribe(ctx, "test-client", query)
	if err != nil {
		panic(err)
	}
	fmt.Println("subscribed to new block events")

	// On NewBlock event
	go func() {
		for data := range rsp {

			// get event data
			events := data.Events

			// process the events
			processEvents(events)
		}
	}()

	select {} // block forever

}
