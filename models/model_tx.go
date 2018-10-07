package models

import (
	"fmt"
	"time"
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

func (tx *TX) Digest() []byte {
	h := sha3.New256()
	h.Write(tx.Serialize())
	return h.Sum(nil)
}

func (tx *TX) Hash() string {
	digest := tx.Digest()
	return fmt.Sprintf("%x", digest)
}

func (tx *TX) Serialize() []byte {

	a := make([][]interface{}, 3)

	a[0] = []interface{}{
		time.Now().UTC().Unix(),
	}

	for _, input := range tx.Inputs {
		a[1] = append(
			a[1],
			input.Serialize(),
		)
	}

	for _, output := range tx.Outputs {
		a[2] = append(
			a[2],
			output.Serialize(),
		)
	}

	b, _ := json.Marshal(a)

	return b
}
