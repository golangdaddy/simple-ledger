package models

import (
	"fmt"
)

func NewOutputIssue(name string, units, maxQuantity int) OutputPayload {
	return OutputIssue{
		Type: "issue",
		Name: name,
		Quantity: maxQuantity,
	}
}

type OutputIssue struct {
	Type string `json:"type"`
	// unique name of token
	Name string `json:"name"`
	// max quantity of the token to issue
	Quantity int `json:"quantity"`
}

func (self OutputIssue) PayloadType() string {
	return self.Type
}

func (self OutputIssue) Serialize() string {
	return fmt.Sprintf("%s:%s:%d", self.Type, self.Name, self.Quantity)
}
