package models

import (
	"fmt"
)

type OutputAsset struct {
	// unique name of token
	Name string
	// value in units of the token
	Value int
}

func (self OutputAsset) Serialize() string {
	return fmt.Sprintf("%s:%d", self.Name, self.Value)
}
