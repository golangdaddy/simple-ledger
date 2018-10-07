package main

import (
	"github.com/golangdaddy/tarantula/web"
	//
//	"github.com/golangdaddy/simple-ledger/models"
)

func (app *App) apiGetInfo(req web.RequestInterface) *web.ResponseStatus {

	app.info.RLock()
	defer app.info.RUnlock()

	return req.Respond(app.info)
}
