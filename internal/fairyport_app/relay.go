package fairyport_app

import (
	"github.com/Fairblock/fairyport/pkg/cosmos/transaction"
)

func RelayDecryptionKeyToCosmos(
	app *FairyportApp,
	height uint64,
	decryptionKey string,
) error {
	err := transaction.SendTx(
		app.CosmosAccountInfo,
		app.TxClient,
		height,
		decryptionKey,
		app.Cfg.CosmosRelayConfig.DestinationNode.ChainId,
	)
	return err
}
