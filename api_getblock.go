package main

import (
	"fmt"
	//
	"github.com/golangdaddy/tarantula/web"
	//
	"github.com/golangdaddy/simple-ledger/models"
)

func (app *App) apiGetBlock(req web.RequestInterface) *web.ResponseStatus {

	req.Log().Debug("CAN YOU SEE MEE")

	blockHeight := req.Param("blockHeight").(int)

	serial := []interface{}{}
	_, err := app.Get(
		fmt.Sprintf("block_%d", blockHeight),
		&serial,
	)
	if req.Log().Error(err) {
		return req.Respond(400, err)
	}

	block := models.ParseBlock(serial)

	return req.Respond(block)
}
