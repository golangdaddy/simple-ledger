package merkletree

func (tree *Tree) digest(b []byte) []byte {
	h := tree.hash()
	h.Write(b)
	return h.Sum(nil)
}
