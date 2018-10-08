package main

import (
	"fmt"
	//
	"github.com/golangdaddy/tarantula/web"
	//
	"github.com/golangdaddy/simple-ledger/models"
)

func (app *App) apiCreateKeypair(req web.RequestInterface) *web.ResponseStatus {
	keypair, err := models.NewAddress()
	if req.Log().Error(err) {
		return req.Fail()
	}
	return req.Respond(
		map[string]string{
			"public": fmt.Sprintf("%x", keypair.Public),
			"private": fmt.Sprintf("%x", keypair.Private),
		},
	)
}
