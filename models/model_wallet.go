package models

import (
	"fmt"
	"crypto/rand"
	"crypto/sha512"
	//
	"github.com/tyler-smith/go-bip32"
)

func NewWallet() *Wallet {
	address, _ := NewAddress()
	return &Wallet{
		Addresses: []*Address{address},
	}
}

type Wallet struct {
	Addresses []*Address
}

type Address struct {
	Public []byte
	Private []byte
}

func (address *Address) Address() string {
	return fmt.Sprintf("%x", address.Public)
}

func NewAddress() (*Address, error) {
	seed := make([]byte, 64)
	rand.Read(seed)
	key, err := keyFromSeed(seed, 10, 1)
	if err != nil {
		return nil, err
	}
	return &Address{
		Public: key.PublicKey().Key,
		Private: key.Key,
	}, err
}

func keyFromSeed(input []byte, difficulty, index int) (*bip32.Key, error) {

	for x := 0; x < difficulty; x++ {

		h := sha512.New512_256()
		h.Write(input)
		input = append(input, h.Sum(nil)...)

	}

	h := sha512.New512_256()
	h.Write(input)
	input = h.Sum(nil)

	masterKey, err := bip32.NewMasterKey(input)
	if err != nil {
		return nil, err
	}

	childKey, err := masterKey.NewChildKey(uint32(index))
	if err != nil {
		return nil, err
	}

	return childKey, nil
}
