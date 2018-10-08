package main

import (
	"github.com/golangdaddy/tarantula/web"
	//
	"github.com/golangdaddy/simple-ledger/models"
)

func (app *App) apiPermissionGrant(req web.RequestInterface) *web.ResponseStatus {

	address := req.BodyParam("address").(string)
	actions := req.BodyParam("actions").([]string)

	tx := &models.TX{
		Inputs: []models.Input{

		},
		Outputs: []models.Output{
			models.Output{
				address,
				models.NewOutputPermission(actions),
			},
		},
	}

	app.SendTransaction(tx)

	return nil
}
