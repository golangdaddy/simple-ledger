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
	duration, err := time.ParseDuration(fmt.Sprintf("%vs", seconds))
	if err != nil {
		panic(err)
	}

	// make genesis block, send it to the block handler
	block := (&models.MainBlock{}).Next(app.wallet.Addresses[0], CONST_COINBASE_REWARD)
	app.blockChannel <- block

	log.Print("Genesis block created..")

	for {
		// create a new block
		block = block.Next(app.wallet.Addresses[0], CONST_COINBASE_REWARD)
		log.Print("Producing new block with height ", block.BlockHeight)

		timer := time.NewTimer(duration)

		for {
			select {

			case <- timer.C:

					log.Print("BLOCK TIME LIMIT REACHED")

					app.blockChannel <- block

				case tx := <- app.txChannel:

					if err := app.parseTransaction(tx); err != nil {
						log.Print(err)
						continue
					}

					block.AddTX(tx)
					continue

			}
			break
		}
	}

}

func (app *App) parseTransaction(tx *models.TX) error {

	app.info.Lock()
		app.info.Index(tx.Txid, tx)
	app.info.Unlock()

	// validate tx
	if len(tx.Outputs) == 0 {
		return fmt.Errorf("THERE ARE NO OUTPUTS ON THIS TRANSACTION")
	}

	for _, output := range tx.Outputs {
		switch output.Payload.PayloadType() {

			// grant /revoke
			case "permission":
				app.info.TotalNativeCurrency += CONST_COINBASE_REWARD

				payload := output.Payload.(models.OutputPermission)
				app.info.Lock()
					app.info.GrantPermissions(output.Recipient, payload.Actions)
					app.DebugJSON(app.info.Permissions(output.Recipient))
				app.info.Unlock()

			case "coinbase":


			default:

				return fmt.Errorf("WTF "+output.Payload.PayloadType())

		}
	}

	return nil
}

func (app *App) blockHandler() {

	log.Print("Waiting for new blocks...")

	for {
		block := <- app.blockChannel

		block.BlockHash = block.Hash()

		log.Print("New Block:")
		log.Print("")
		app.DebugJSON(block)
		log.Print("")

		if err := app.PutBlock(block); err != nil {
			panic(err)
		}

		// update the chain info
		app.info.Lock()
			app.info.TotalTransactions += block.TxCount
			app.info.BlockHeight = block.BlockHeight
			app.info.BlockHash = block.BlockHash
		app.info.Unlock()

		log.Print("Written to database!")
		log.Print("")
	}

}

func (app *App) loadChain() {
	log.Print("Loading chain...")
}
