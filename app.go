package main

import (
	"os"
	"fmt"
	"log"
	"encoding/json"
	//
	"github.com/dgraph-io/badger"
	"github.com/golangdaddy/simple-ledger/models"
)

type App struct {
	badgerDB *badger.DB
	chainName string
	info *ChainInfo
	wallet *models.Wallet
	params map[string]interface{}
	txChannel chan *models.TX
	blockChannel chan *models.MainBlock
}

// Filesystem is for creating safe paths when accessing the filesystem
func (app *App) Filesystem(path string) string {
	chainData := fmt.Sprintf("chaindata/%s", app.chainName)
	log.Print("Creating data folder: "+chainData)
	os.Mkdir(chainData, 0777)
	return fmt.Sprintf("%s/%s", chainData, path)
}

func (app *App) DebugJSON(x interface{}) {
	b, _ := json.Marshal(x)
	log.Print(string(b))
}
