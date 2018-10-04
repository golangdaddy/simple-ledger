package models

import (
	"fmt"
)

// Defines an interface for instructing the output to do something
type OutputPayload interface {
	Serialize() string
}

type Output struct {
	// address receiving the token
	Recipient string
	// type of thing, e.g. asset, or permission
	Payload OutputPayload
}

func (self Output) Serialize() string {
	return fmt.Sprintf("%s@%s", self.Recipient, self.Payload.Serialize())
}
