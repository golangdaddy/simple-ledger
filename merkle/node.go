package merkletree

import (
	"bytes"
)

type Leaves []*Node

func (self Leaves) Len() int {
	return len(self)
}

func (self Leaves) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

func (self Leaves) Less(i, j int) bool {
	return bytes.Compare(self[i].Hash, self[j].Hash) == -1
}

type Node struct {
	Raw []byte
	Hash []byte
	Size int
	Decendants int
	children []*Node
}

func (tree *Tree) Digest(node *Node) []byte {

	if node.children == nil {
		return node.Hash
	}

	childCount := len(node.children)
	subHashes := make([][]byte, childCount)

	for x, child := range node.children {
		subHashes[x] = tree.Digest(child)
	}
	return tree.digest(
		bytes.Join(subHashes, nil),
	)
}

func (tree *Tree) newLeaf(b []byte) *Node {
	return &Node{
		Raw: b,
		Hash: tree.digest(b),
		Size: len(b),
		Decendants: 1,
	}
}
