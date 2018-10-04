package main

import (
	"os"
	"io/ioutil"
	"crypto/rand"
	"crypto/x509"
	"crypto/ecdsa"
	"crypto/elliptic"
)

const (
	CONST_PATH_KEYS = "keys/"
	CONST_PATH_PUBLIC_KEY = CONST_PATH_KEYS + "public.key"
	CONST_PATH_PRIVATE_KEY = CONST_PATH_KEYS + "private.key"
)

func (app *App) NewKeys() (b []byte, privateKey *ecdsa.PrivateKey, err error) {

	privateKey, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}

	os.Mkdir("keys", 0777)

	// Write private key
	b, err = x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		return nil, nil, err
	}
	err = ioutil.WriteFile(CONST_PATH_PRIVATE_KEY, b, 0666)
	if err != nil {
		return nil, nil, err
	}

	// Write public key
	b, err = x509.MarshalPKIXPublicKey(privateKey)
	if err != nil {
		return nil, nil, err
	}
	err = ioutil.WriteFile(CONST_PATH_PUBLIC_KEY, b, 0666)
	if err != nil {
		return nil, nil, err
	}

	return b, privateKey, nil
}
