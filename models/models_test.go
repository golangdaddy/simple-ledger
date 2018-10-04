package models

import (
	"fmt"
	"testing"
)

func newTX() *TX {
	return &TX{
		Inputs: []Input{
			Input{
				Txid: "1234",
				Index: 0,
			},
		},
		Outputs: []Output{},
	}
}

func TestSerialization(t *testing.T) {

	t.Run(
		"Check asset output",
		func (t *testing.T) {

			tx := newTX()
			tx.Output(
				"abdc",
				OutputAsset{
					Name: "BTC",
					Value: 1000,
				},
			)

			serial := tx.Serialize()
			fmt.Println(string(serial[:]), "==", `[["1234:0"],["abdc@BTC:1000"]]`)

		},
	)

	t.Run(
		"Check issue output",
		func (t *testing.T) {

			tx := newTX()
			tx.Output(
				"abdc",
				OutputIssue{
					Name: "BTC",
					Units: 100,
					Quantity: 10000,
				},
			)

			serial := tx.Serialize()
			fmt.Println(string(serial[:]), "==", `[["1234:0"],["abdc@BTC:1000"]]`)

		},
	)

	t.Run(
		"Check permission output",
		func (t *testing.T) {

			tx := newTX()
			tx.Output(
				"abdc",
				OutputPermission{
					Actions: []string{
						"send",
						"receive",
					},
				},
			)

			serial := tx.Serialize()
			fmt.Println(string(serial[:]), "==", `[["1234:0"],["abdc@BTC:1000"]]`)

		},
	)

}
