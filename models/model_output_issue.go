package models

import (
	"fmt"
)

type OutputIssue struct {
	// unique name of token
	Name string
	// divisible units of the token
	Units int
	// quantity of the token to issue
	Quantity int
	// whether quantities of the asset can be issued in the future
	Open bool
}

func (self OutputIssue) Serialize() string {
	return fmt.Sprintf("%s:%d:%d:%v", self.Name, self.Units, self.Quantity, self.Open)
}
