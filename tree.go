/*
 Taken from
 http://cslibrary.stanford.edu/110/BinaryTrees.html
*/

package main

import (
	//"fmt"
	"log"
)

func main() {
	log.Print("Building tree")
	Build123()
}

func Build123() *Node {
	root := NewNode(2)
	lchild := NewNode(1)
	rchild := NewNode(3)
	root.left = lchild
	root.right = rchild

	return root
}

// NOT WORKING FFS
type Node struct {
	left  *Node
	right *Node
	data  int
}

type Compare func(i, j int) int8

/*
 Given a binary tree, return true if a node
 with the target data is found in the tree. Recurs
 down the tree, chooses the left or right
 branch by comparing the target to each node.
*/
func Lookup(n Node, target int) int {
	// 1. Base case == empty tree
	// in that case, the target is not found so return false
	if n == nil {
		return -1
	} else {
		// 2. see if found here
		if target == n.data {
			return n.data
		} else {
			// 3. otherwise recur down the correct subtree
			if target < n.GetData() {
				return Lookup(n.left, target)
			} else {
				return Lookup(n.right, target)
			}
		}
	}
}

/*
 Helper function that allocates a new node
 with the given data and NULL left and right
 pointers.
*/
func NewNode(data int) *Node {
	n := new(&Node)
	n.data = data
	//n.left = nil
	//n.right = nil
	return n
}

/*
 Give a binary search tree and a number, inserts a new node
 with the given number in the correct place in the tree.
 Returns the new root pointer which the caller should
 then use (the standard trick to avoid using reference
 parameters).
*/
func Insert(n Node, data int) Node {
	// 1. If the tree is empty, return a new, single node
	if n == nil {
		return NewNode(data)
	} else {
		// 2. Otherwise, recur down the tree
		if data <= n.GetData() {
			n.left = Insert(n.left, data)
		} else {
			n.right = Insert(n.right, data)
		}
		return n
	}
}
