package main

import (
	"log"
	"fmt"
	"time"
	//
	"github.com/golangdaddy/simple-ledger/models"
)

func (app *App) initChain() {
	log.Print("Creating chain...")

	go app.blockHandler()

	// read the blocktime param
	seconds, err := app.IntParam("blocktime")
	if err != nil {
		panic(err)
	}
	duration, err := time.ParseDuration(fmt.Sprintf("%ds", seconds))
	if err != nil {
		panic(err)
	}

	// make genesis block, send it to the block handler
	block := (&models.MainBlock{}).Next(app.wallet.Addresses[0], CONST_COINBASE_REWARD)
	app.blockChannel <- block

	for {
		// create a new block
		block = block.Next(app.wallet.Addresses[0], CONST_COINBASE_REWARD)

		for {
			select {

				case tx := <- app.txChannel:

					app.info.Index(tx.Txid, tx)

					// validate tx
					if len(tx.Outputs) == 0 {
						log.Print("THERE ARE NO OUTPUTS ON THIS TRANSACTION")
						continue
					}

					app.DebugJSON(tx)

					for _, output := range tx.Outputs {
						switch output.Payload.PayloadType() {
							case "permission":
								app.info.TotalNativeCurrency += CONST_COINBASE_REWARD

								payload := output.Payload.(models.OutputPermission)
								app.info.GrantPermissions(output.Recipient, payload.Actions)
								app.DebugJSON(app.info.Permissions(output.Recipient))

							case "coinbase":


							default:

								log.Print("WTF "+output.Payload.PayloadType())

						}
					}

					block.AddTX(tx)
					continue

				case <- time.After(duration):

					app.blockChannel <- block

			}
			break
		}
	}

}

func (app *App) blockHandler() {

	for {
		block := <- app.blockChannel

		block.BlockHash = block.Hash()
		serial := block.Serial()

		app.DebugJSON(block)

		// update the chain info
		app.info.Lock()
			app.info.TotalTransactions += block.TxCount
			app.info.BlockHeight = block.BlockHeight
			app.info.BlockHash = block.BlockHash
		app.info.Unlock()

		if err := app.Update(block.BlockHash, serial); err != nil {
			panic(err)
		}
	}

}

func (app *App) loadChain() {
	log.Print("Loading chain...")
}
