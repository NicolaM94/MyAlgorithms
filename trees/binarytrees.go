package trees



type Node struct {
	Value int
	Left *Node
	Right *Node
}

func (n *Node) AddLeftChild(value int) {
	n.Left.Value = value
}

func (n *Node) AddRightChild (value int) {
	n.Right.Value = value
}

type BinaryTree struct {
	Root *Node
}


