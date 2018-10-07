package main

import (
	"fmt"
	"reflect"
)

func (app *App) StringParam(k string) (string, error) {

	x, ok := app.params[k].(string)
	if !ok {
		return "", fmt.Errorf("Failed to type assert %s to string!", reflect.TypeOf(app.params[k]).String())
	}

	return x, nil
}

func (app *App) IntParam(k string) (int, error) {

	x, ok := app.params[k].(int)
	if !ok {
		if app.params[k] == nil {
			return 0, fmt.Errorf("Failed to type assert NIL to int!")
		}
		return 0, fmt.Errorf("Failed to type assert %s to int!", reflect.TypeOf(app.params[k]).String())
	}

	return x, nil
}
