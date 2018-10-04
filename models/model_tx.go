package models

import (
	"encoding/json"
	//
	"golang.org/x/crypto/sha3"
)

type TX struct {
	// digest of this transaction
	Txid string
	// where the tokens come from
	Inputs []Input
	// where the tokens go
	Outputs []Output
}

func (self *TX) Output(recipient string, output OutputPayload) {
	self.Outputs = append(
		self.Outputs,
		Output{
			Recipient: recipient,
			Payload: output,
		},
	)
}

func (tx *TX) Digest() [32]byte {
	return sha3.Sum256(tx.Serialize())
}

func (tx *TX) Serialize() []byte {

	a := make([][]interface{}, 2)

	for _, input := range tx.Inputs {
		a[0] = append(
			a[0],
			input.Serialize(),
		)
	}

	for _, output := range tx.Outputs {
		a[1] = append(
			a[1],
			output.Serialize(),
		)
	}

	b, _ := json.Marshal(a)

	return b
}
