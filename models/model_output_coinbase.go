package models

import (
	"fmt"
)

func NewOutputCoinbase(reward float64) OutputPayload {
	return OutputCoinbase{
		Type: "coinbase",
		NativeReward: reward,
	}
}

type OutputCoinbase struct {
	// unique name of token
	Type string `json:"type"`
	NativeReward float64 `json:"nativeReward"`
}

func (self OutputCoinbase) PayloadType() string {
	return self.Type
}

func (self OutputCoinbase) Serialize() string {
	return fmt.Sprintf("%s:%d", self.Type, self.NativeReward)
}
