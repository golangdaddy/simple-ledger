package models

import (
	"fmt"
)

func NewOutputMint(name string, units, maxQuantity int) OutputPayload {
	return OutputMint{
		Type: "mint",
		Name: name,
		Units: units,
		MaxQuantity: maxQuantity,
	}
}

type OutputMint struct {
	Type string `json:"type"`
	// unique name of token
	Name string `json:"name"`
	// divisible units of the token
	Units int `json:"units"`
	// max quantity of the token to issue
	MaxQuantity int `json:"maxQuantity"`
}

func (self OutputMint) PayloadType() string {
	return self.Type
}

func (self OutputMint) Serialize() string {
	return fmt.Sprintf("%s:%s:%d:%d", self.Type, self.Name, self.Units, self.MaxQuantity)
}
