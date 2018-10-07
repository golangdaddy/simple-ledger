package main

import (
	"fmt"
	"net/http"
	//
	"github.com/golangdaddy/tarantula/log/testing"
	"github.com/golangdaddy/tarantula/router/common"
	"github.com/golangdaddy/tarantula/router/standard"
	//
	"github.com/golangdaddy/simple-ledger/models"
)

func (app *App) httpServer() {

	logClient := logs.NewClient().NewLogger()

	root, service := router.NewRouter(logClient, CONST_SERVICENAME)

	root.Config.SetHeaders(
		common.Headers{
			"Access-Control-Allow-Headers": "Authorization,Content-Type",
			"Access-Control-Allow-Origin": "*",
			"Access-Control-Allow-Methods": "PUT, POST, GET, DELETE, OPTIONS",
		},
	)

	root.Add("/sdjfjhj").POST(
		app.apiGetInfo,
	).Describe(
		"Gets the info about the chain, or it gets the hose again!",
	).Response(
		models.GetInfo{},
	)

	rpcport, err := app.IntParam("rpcport")
	if err != nil {
		panic(err)
	}

	panic(
		http.ListenAndServe(
			fmt.Sprintf(":%d", rpcport),
			service,
		),
	)
}
