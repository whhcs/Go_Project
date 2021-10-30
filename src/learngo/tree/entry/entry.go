package main

import (
	"fmt"

	"github.com/whhcs/learngo/tree"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	left.postOrder()

	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()
}

func main() {
	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	/*	nodes := []treeNode{
			{value: 3},
			{},
			{6, nil, &root},
		}
		fmt.Println(nodes)*/
	root.Right.Left.SetValue(4)

	root.Traverse()
	fmt.Println()

	MyRoot := myTreeNode{&root}
	MyRoot.postOrder()
	fmt.Println()
}
