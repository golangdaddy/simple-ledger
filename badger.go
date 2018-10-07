package main

import (
	"github.com/dgraph-io/badger"
)

func (app *App) initBadger() {
	// Open the Badger database located in the /tmp/badger directory.
	// It will be created if it doesn't exist.

	dataPath := app.Filesystem("badger")

	opts := badger.DefaultOptions
	opts.Dir = dataPath
	opts.ValueDir = dataPath
	db, err := badger.Open(opts)
	if err != nil {
		panic(err)
	}

	app.badgerDB = db
}
