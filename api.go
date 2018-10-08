package main

import (
	"fmt"
	"net/http"
	//
	"github.com/golangdaddy/tarantula/web/validation"
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

	root.Add("/info").GET(
		app.apiGetInfo,
	).Describe(
		"Gets the info about the chain, or it gets the hose again!",
	).Response(
		ChainInfo{},
	)

	root.Add("/block").Param(validation.Int(), "blockHeight").GET(
		app.apiGetBlock,
	).Describe(
		"Gets the block at the specified height.",
	).Response(
		models.MainBlock{},
	)

	permission := root.Add("/permission")

		permission.Add("/grant").GET(
			app.apiPermissionGrant,
		).Describe(
			"Grants specified actions to the given address.",
		).Body(
			&common.Payload{
				"address": validation.Hex256(),
				"actions": validation.ArrayString(),
			},
		)

	create := root.Add("/create")

		create.Add("/keypair").GET(
			app.apiCreateKeypair,
		).Describe(
			"Creates a new keypair that is added to the wallet.",
		).Response(
			models.Address{},
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
