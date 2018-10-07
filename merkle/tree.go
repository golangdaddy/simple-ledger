package merkletree

import (
	"sync"
	"sort"
	"hash"
)

type Tree struct {
	hash func () hash.Hash
	index map[string]*Node
	matrix [][]*Node
	root *Node
	sync.RWMutex
}

func New(h func () hash.Hash) *Tree {
	return &Tree{
		index: map[string]*Node{},
		root: &Node{},
		matrix: [][]*Node{
			[]*Node{},
			[]*Node{},
			[]*Node{},
			[]*Node{},
			[]*Node{},
			[]*Node{},
			[]*Node{},
			[]*Node{},
			[]*Node{},
			[]*Node{},
			[]*Node{},
			[]*Node{},
			[]*Node{},
			[]*Node{},
			[]*Node{},
			[]*Node{},
			[]*Node{},
			[]*Node{},

		},
		hash: h,
	}
}

func (tree *Tree) Add(k string, b []byte) {

	leaf := tree.newLeaf(b)

	tree.Lock()

		tree.matrix[0] = append(
			tree.matrix[0],
			leaf,
		)

		tree.index[k] = leaf

		sort.Sort(Leaves(tree.matrix[0]))

		tree.build()

	tree.Unlock()

}

func (tree *Tree) build() {

	depth := 0
	lastNode := tree.matrix[depth][0]

	for len(tree.matrix[depth]) > 1 {

		nodes := []*Node{}

		for x, node := range tree.matrix[depth] {
			// create a new block every 2 nodes
			if x % 2 == 0 {
				lastNode = &Node{}
				nodes = append(
					nodes,
					lastNode,
				)
			}
			lastNode.Decendants += node.Decendants
			lastNode.children = append(lastNode.children, node)
		}

		depth++
		tree.matrix[depth] = nodes

	}

	tree.root = lastNode

}

func (tree *Tree) Root() []byte {

	tree.RLock()
	defer tree.RUnlock()

	return tree.Digest(tree.root)
}
