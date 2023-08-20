package trees

// Node structure to use in a tree
type Node struct {
	Value int
	Left *Node
	Right *Node
}

// Adds node (left child) to another node (parent)
func (n *Node) AddLeftChild(value int) {
	n.Left.Value = value
}

// Adds node (right child) to another node (parent)
func (n *Node) AddRightChild (value int) {
	n.Right.Value = value
}

// Structure containing the tree, with its root node
type BinaryTree struct {
	Root *Node
}


