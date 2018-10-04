package main

import (
	"os"
	"fmt"
	"log"
	"strings"
	//
	"github.com/dgraph-io/badger"
)

type App struct {
	badgerDB *badger.DB
	chainName string
}

func main() {

	app := &App{}

	args := os.Args
	if len(args) != 2 {
		log.Fatal("No chain name arguement provided!")
		os.Exit(10)
	}
	app.chainName = args[1]

	fmt.Println("Using example chain with blockchain name: "+app.chainName)

	_, err := os.Stat("chaindata/"+app.chainName)
	if err != nil {
		if strings.Contains(
			err.Error(),
			"no such file or directory",
		) {
			fmt.Printf("Would you like to create a new blockchain called '%s'? Y/n\n", app.chainName)
		}
	}

	app.initBadger()


	<- make(chan bool)

}
