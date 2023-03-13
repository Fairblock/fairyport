package fairyclient

import (
	"context"
	"log"
	"time"

	"github.com/FairBlock/fairyport/internal/events"
	"github.com/FairBlock/fairyport/pkg/account"
	"github.com/FairBlock/fairyport/pkg/transaction"
	"github.com/cosmos/cosmos-sdk/types/tx"

	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
)

func StartFairyClient(fairyClient *rpchttp.HTTP, accDetails *account.AccountDetails, txClient tx.ServiceClient) {
	err := fairyClient.Start()
	if err != nil {
		log.Fatal(err)
	}

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
	log.Println("subscribed to new block events")

	// On NewBlock event
	go func() {
		for data := range rsp {

			// get event data
			blockEvents := data.Events

			// process the events
			height, data := events.ProcessEvents(blockEvents)

			err := transaction.SendTx(accDetails, txClient, height, data)
			if err != nil {
				log.Println("Sending Transaction for height :", height, " failed")
			}

		}
	}()

	select {} // block forever
}
