package main

import (
	"os"
	"fmt"
	"log"
	"bufio"
	"strings"
	"io/ioutil"
)

const (
	CONST_SERVICENAME = "SimpleLedger"
	CONST_COINBASE_REWARD = 10.0
)

func main() {

	os.Mkdir("chaindata", 0777)

	args := os.Args
	if len(args) != 2 {
		log.Fatal("No chain name arguement provided!")
		os.Exit(10)
	}

	app, err := initApp(args[1])
	if err != nil {
		panic(err)
	}

	fmt.Println("Using example chain with blockchain name: "+app.chainName)

	_, err = os.Stat(app.Filesystem("params.dat"))
	if err != nil {
		if strings.Contains(
			err.Error(),
			"no such file or directory",
		) {
			for {
				reader := bufio.NewReader(os.Stdin)
				fmt.Printf("Would you like to create a new blockchain called '%s'? Y/n\n", app.chainName)
				text, _ := reader.ReadString('\n')
				input := strings.ToLower(string(text[0]))

				if input == "y" {
					break
				}
				if input == "n" {
					fmt.Println("Shutting down application...")
					os.Exit(10)
				}

			}

		} else {
			panic(err)
		}
	}

	// write parameters to disk
	err = ioutil.WriteFile(app.Filesystem("params.dat"), []byte(""), 0777)
	if err != nil {
		panic(err)
	}

	app.initChain()

	<- make(chan bool)

}
