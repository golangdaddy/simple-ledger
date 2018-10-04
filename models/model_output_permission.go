package models

import (
	"fmt"
	"strings"
)

type OutputPermission struct {
	// unique name of token
	Actions []string
}

func (self OutputPermission) Serialize() string {
	return fmt.Sprintf("%s", strings.Join(self.Actions, ","))
}
