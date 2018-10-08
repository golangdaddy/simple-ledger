package models

import (
	"fmt"
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
	Transactions []string `json:"transactions"`
	transactions []*TX
	fromDisk bool
	tree *merkletree.Tree
}

func (block *MainBlock) FromDisk() bool {
	return block.fromDisk
}

func ParseBlock(serial []interface{}) *MainBlock {

	block := &MainBlock{
		fromDisk: true,
		BlockHeight: int(serial[0].(float64)),
		Timestamp: int64(serial[1].(float64)),
		PrevBlockHash: serial[2].(string),
		TxCount: int(serial[3].(float64)),
		MerkleRoot: serial[4].(string),
		//Transactions = serial[5].([]interface{})
	}
	block.Hash()
	return block
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
		tree: merkletree.New(sha3.New256),
		Timestamp: time.Now().UTC().Unix(),
		BlockHeight: block.BlockHeight + 1,
		PrevBlockHash: block.BlockHash,
	}
	nextBlock.AddCoinbase(address, reward)
	return nextBlock
}

func (block *MainBlock) AddTX(tx *TX) {

	//log.Print("Including transaction in block: "+tx.Hash())

	block.transactions = append(
		block.transactions,
		tx,
	)

	for _, tx := range block.transactions {
		block.tree.Add(
			"",
			tx.Serialize(),
		)
	}

	// update block details
	block.TxCount++
	block.Timestamp = time.Now().UTC().Unix()

}

func (block *MainBlock) Hash() {

	if !block.fromDisk {
		block.MerkleRoot = fmt.Sprintf(
			"%x",
			block.tree.Root(),
		)
	}

	b, _ := json.Marshal(block.HeaderArray())

	for x := 0; x < 2; x++ {
		h := sha3.New256()
		h.Write(b)
		b = h.Sum(nil)
	}

	block.BlockHash = fmt.Sprintf("%x", b)
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
	for x, tx := range block.transactions {
		block.Transactions[x] = tx.Hash()
	}

	serial := append(
		block.HeaderArray(),
		block.BlockHash,
		block.Transactions,
	)

	b, _ := json.Marshal(serial)

	return b
}
