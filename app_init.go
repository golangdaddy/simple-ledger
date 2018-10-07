package main

import (
	"github.com/dgraph-io/badger"
	"github.com/golangdaddy/simple-ledger/models"
)


// init creates the initial app state
func initApp(chainName string) (app *App, err error) {

	app = &App{
		chainName: chainName,
		wallet: models.NewWallet(),
		params: map[string]interface{}{},
		txChannel: make(chan *models.TX),
		blockChannel: make(chan *models.MainBlock),
		info: &models.GetInfo{},
	}

	// temp
	app.params["blocktime"] = 5
	app.params["rpcport"] = 6789

	dataPath := app.Filesystem("badger")

	opts := badger.DefaultOptions
	opts.Dir = dataPath
	opts.ValueDir = dataPath
	app.badgerDB, err = badger.Open(opts)
	if err != nil {
		return nil, err
	}

	go app.debugTranactions()

	go app.httpServer()

	return app, nil
}
