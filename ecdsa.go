package main

import (
	"crypto/rand"
	"crypto/ecdsa"
	"crypto/elliptic"
)

const (
	CONST_PATH_KEYS = "keys/"
	CONST_PATH_PUBLIC_KEY = CONST_PATH_KEYS + "public.key"
	CONST_PATH_PRIVATE_KEY = CONST_PATH_KEYS + "private.key"
)

func (app *App) NewKeys() (privateKey *ecdsa.PrivateKey, err error) {

	privateKey, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}
