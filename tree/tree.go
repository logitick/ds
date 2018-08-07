package tree

type Tree interface {
	Insert(v interface{})
	Delete(v interface{}) interface{}
	Search(v interface{})
	IsEmpty() bool
	Size() int  // counts the number of nodes in the tree
	Depth() int // counts the number of nodes along the longest path from the root node down to the farthest leaf node
}

// Node a tree node
type Node struct {
	value interface{}
	left  *Node
	right *Node
}

// BinarySearchTree implements tree specified in http://cslibrary.stanford.edu/110/BinaryTrees.html
type BinarySearchTree struct {
	root *Node
}

// New creates a new instance of a BST
func New() *BinarySearchTree {
	return &BinarySearchTree{root: nil}
}

// Size counts the number of nodes inside the tree
func (t *BinarySearchTree) Size() int {
	return 0
}

// Insert wraps the value in a Node and adds it to the tree
func (t *BinarySearchTree) Insert(v interface{}) {

}
