package models

import (
	"fmt"
	"strings"
)

func NewOutputPermission(actions []string) OutputPayload {
	return OutputPermission{
		Type: "permission",
		Actions: actions,
	}
}

type OutputPermission struct {
	Type string `json:"type"`
	// unique name of token
	Actions []string `json:"actions"`
}

func (self OutputPermission) PayloadType() string {
	return self.Type
}

func (self OutputPermission) Serialize() string {
	return fmt.Sprintf("%s:%s", self.Type, strings.Join(self.Actions, ","))
}
