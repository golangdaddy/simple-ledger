package models

import (
	"fmt"
)

type Input struct {
	// transaction being referenced
	Txid string
	// output index within referenced transaction
	Index int
}

func (self Input) Serialize() string {
	return fmt.Sprintf(
		"%s:%d",
		self.Txid,
		self.Index,
	)
}
