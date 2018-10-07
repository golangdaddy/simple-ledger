package models

import (
	"fmt"
	"log"
	"time"
	"encoding/json"
	//
	"golang.org/x/crypto/sha3"
	"github.com/golangdaddy/simple-ledger/merkle"
)

type MainBlock struct {
	Timestamp int64 `json:"timestamp"`
	BlockHash string `json:"hash,omitempty"`
	BlockHeight int `json:"height"`
	PrevBlockHash string `json:"prevBlockHash"`
	MerkleRoot string `json:"merkleRoot"`
	TxCount int `json:"txCount"`
	transactions []*TX `json:"-"`
	Transactions []string `json:"transactions"`
}

func (block *MainBlock) AddCoinbase(address *Address, reward float64) {
	block.AddTX(
		&TX{
			Inputs: []Input{},
			Outputs: []Output{
				Output{
					address.Address(),
					NewOutputCoinbase(reward),
				},
			},
		},
	)
}

func (block *MainBlock) Next(address *Address, reward float64) *MainBlock {
	nextBlock := &MainBlock{
		BlockHeight: block.BlockHeight + 1,
		PrevBlockHash: block.Hash(),
	}
	nextBlock.AddCoinbase(address, reward)
	return nextBlock
}

func (block *MainBlock) AddTX(tx *TX) {

	log.Print("Including transaction in block: "+tx.Hash())

	block.transactions = append(
		block.transactions,
		tx,
	)

	tree := merkletree.New(sha3.New256)

	for _, tx := range block.transactions {
		tree.Add(
			"",
			tx.Serialize(),
		)
	}

	// update block details
	block.TxCount++
	block.Timestamp = time.Now().UTC().Unix()
	block.MerkleRoot = fmt.Sprintf(
		"%x",
		tree.Root(),
	)

}

func (block *MainBlock) Hash() string {

	b, _ := json.Marshal(block.HeaderArray())

	for x := 0; x < 2; x++ {
		h := sha3.New256()
		h.Write(b)
		b = h.Sum(nil)
	}

	return fmt.Sprintf("%x", b)
}

func (block *MainBlock) HeaderArray() []interface{} {

	return []interface{}{
		block.BlockHeight,
		block.Timestamp,
		block.PrevBlockHash,
		block.TxCount,
		block.MerkleRoot,
	}
}

func (block *MainBlock) Serial() []byte {

	block.Transactions = make([]string, len(block.transactions))
	for _, tx := range block.transactions {
		block.Transactions = append(
			block.Transactions,
			tx.Hash(),
		)
	}

	b, _ := json.Marshal(
		append(
			block.HeaderArray(),
			block.Transactions,
		),
	)

	return b
}
