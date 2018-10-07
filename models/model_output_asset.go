package models

import (
	"fmt"
)

func NewOutputAsset() OutputPayload {
	return OutputAsset{
		Type: "asset",
	}
}

type OutputAsset struct {
	Type string `json:"type"`
	// unique name of token
	Name string `json:"name"`
	// value in units of the token
	Value int `json:"value"`
}

func (self OutputAsset) PayloadType() string {
	return self.Type
}

func (self OutputAsset) Serialize() string {
	return fmt.Sprintf("%s:%s:%d", self.Type, self.Name, self.Value)
}
