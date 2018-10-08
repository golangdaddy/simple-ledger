package main

import (
	"log"
	"fmt"
	"time"
	//
	"github.com/golangdaddy/simple-ledger/models"
)

func (app *App) rescanChain() (block *models.MainBlock) {

	b, err := app.Get("block_0", nil)
	if err != nil {
		log.Print("Failed to locate Genesis block!")
		return
	}

	lastHash := string(b)
	app.chainUID = lastHash

	log.Print("Scanning chain with genesis ID: "+lastHash)

	for x := 1; x == x; x++ {

		k := fmt.Sprintf("block_%v", x)

		log.Print("Getting block: "+k)

		serial := []interface{}{}
		_, err := app.Get(
			k,
			&serial,
		)
		if err != nil {
			log.Print("Failed to locate: "+k)
			break
		}

		block = models.ParseBlock(serial)

		if lastHash != block.PrevBlockHash {
			log.Print("BLOCKCHAIN IS BROKEN")
			log.Print("'"+lastHash+"' != '"+block.PrevBlockHash+"'")
			return nil
		}

		lastHash = block.BlockHash

		app.blockChannel <- block

	}

	return block
}

func (app *App) initChain() {
	log.Printf("Creating chain %s...", app.chainName)

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

	block := app.rescanChain()
	if block == nil {
		seed := "myblockchainseed"
		log.Print("Creating Genesis block...")
		// make genesis block, send it to the block handler
		block = &models.MainBlock{
			BlockHash: seed,
		}
		block = block.Next(app.wallet.Addresses[0], CONST_COINBASE_REWARD)
		if err := app.Update("block_0", []byte(block.PrevBlockHash)); err != nil {
			panic(err)
		}
		block.Hash()
		app.blockChannel <- block
	}

	log.Print("Genesis block created...")

	var prevBlockHash string

	for {
		prevBlockHash = block.BlockHash

		// create a new block
		block = block.Next(app.wallet.Addresses[0], CONST_COINBASE_REWARD)
		log.Print("Producing new block with height ", block.BlockHeight)

		if prevBlockHash != block.PrevBlockHash {
			panic("INVALID")
		}

		timer := time.NewTimer(duration)

		for {
			select {

			case <- timer.C:

					block.Hash()
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
					//app.DebugJSON(app.info.Permissions(output.Recipient))
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

		log.Print("New Block:")
		log.Print("")
		app.DebugJSON(block)
		log.Print("")

		// only write to disk if this block is new
		if err := app.PutBlock(block); err != nil {
			panic(err)
		}

		// update the chain info
		app.info.Lock()
			app.info.TotalTransactions += block.TxCount
			app.info.BlockHeight = block.BlockHeight
			app.info.BlockHash = block.BlockHash
		app.info.Unlock()

//		log.Print("Written to database!")
	}

}

func (app *App) loadChain() {
	log.Print("Loading chain...")
}
