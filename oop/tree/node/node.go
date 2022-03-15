package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func NewNode(value int) *Node {
	return &Node{
		Value: value,
	}
}

func (n *Node) Print() {
	fmt.Print(n.Value, " ")
}

func (n *Node) SetValue(value int) {
	if n == nil {
		fmt.Println("Setting value to nil node. Ignored.")
		return
	}
	n.Value = value
}
