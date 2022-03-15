package main

import (
	"fmt"
	tree "learning_go/oop/tree/node"
)

// type MyNode struct {
// 	node *tree.Node
// }

// func (myNode *MyNode) postOrder() {
// 	if myNode == nil || myNode.node == nil {
// 		return
// 	}
// 	left := MyNode{myNode.node.Left}
// 	right := MyNode{myNode.node.Right}
// 	left.postOrder()
// 	right.postOrder()
// 	myNode.node.Print()
// }

func main() {
	root := tree.Node{}
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{Value: 5, Left: nil, Right: nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.NewNode(2)
	root.Right.Left.SetValue(4)
	root.Traverse()
	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println("Node count:", nodeCount)
}
