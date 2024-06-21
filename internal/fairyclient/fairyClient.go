package fairyclient

import (
	"context"
	"fmt"
	"github.com/Fairblock/fairyport/config"
	"github.com/Fairblock/fairyport/internal/events"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"time"

	"github.com/Fairblock/fairyport/pkg/account"
	"github.com/Fairblock/fairyport/pkg/transaction"
	"github.com/cosmos/cosmos-sdk/types/tx"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
)

var (
	invalidBroadcastAggregatedKeyShare = promauto.NewCounter(prometheus.CounterOpts{
		Name: "fairyport_invalid_broadcast_aggregated_keyshare",
		Help: "The total number of invalid key share generated",
	})

	validBroadcastAggregatedKeyShare = promauto.NewCounter(prometheus.CounterOpts{
		Name: "fairyport_valid_broadcast_aggregated_keyshare",
		Help: "The total number of valid key share generated",
	})
)

func StartFairyClient(fairyClient *rpchttp.HTTP, accDetails *account.AccountDetails, txClient tx.ServiceClient, cfg config.Config) {
	err := fairyClient.Start()
	if err != nil {
		log.Fatal(err)
	}

	defer fairyClient.Stop()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// query for new block event
	query := "tm.event = 'Tx'"

	// subscribe to new blocks
	rsp, err := fairyClient.Subscribe(ctx, "test-client", query)
	if err != nil {
		panic(err)
	}
	log.Println("subscribed to new block events")
        
	http.Handle("/metrics", promhttp.Handler())
	log.Printf("MetricsPort: %d\n", cfg.MetricsPort)
	go http.ListenAndServe(fmt.Sprintf(":%d", cfg.MetricsPort), nil)

	for data := range rsp {

		// get event data
		blockEvents := data.Events

		// process the events
		height, aggregatedKeyShare, err := events.ProcessEvents(blockEvents)
		if err != nil {
			continue
		}

		err = transaction.SendTx(accDetails, txClient, height, aggregatedKeyShare, cfg)
		if err != nil {
			invalidBroadcastAggregatedKeyShare.Inc()
			log.Println("Sending Transaction for height :", height, " failed: ", err)
			continue
		}
			validBroadcastAggregatedKeyShare.Inc()
		log.Println("Successfully Broadcast Aggregated KeyShare for Height: ", height)
	}
}
