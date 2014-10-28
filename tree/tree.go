package main

import (
	"fmt"
)

type Node struct {
	value               int
	left, right, parent *Node
}

type Tree struct {
	root Node
}

func main() {
	fmt.Println("> Initialize the tree structure")
}

func (n *Node) IsLeaf() bool {
	if n.left = nil && n.right = nil {
		return true
	}
	return false
}

func (t *Tree) Add(n Node) {

}
