package main

import (
	"time"
	//
	"github.com/golangdaddy/simple-ledger/models"
)

func (app *App) debugTranactions() {

	address := app.wallet.Addresses[0]

	for {
		time.Sleep(2 * time.Second)

		app.info.RLock()
			app.DebugJSON(app.info)
		app.info.RUnlock()

		tx := &models.TX{
			Inputs: []models.Input{

			},
			Outputs: []models.Output{
				models.Output{
					address.Address(),
					models.NewOutputPermission([]string{"send", "receive"}),
				},
			},
		}

		tx.Txid = tx.Hash()
		app.txChannel <- tx
	}
}
