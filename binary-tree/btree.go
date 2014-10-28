/*
 Taken from
 http://cslibrary.stanford.edu/110/BinaryTrees.html
*/

package tree

import (
	//"fmt"
	"log"
)

func main() {
	log.Print("Building tree")
	Build123()
}

func Build123() (b *Btree) {
	root := NewNode(2)
	lchild := NewNode(1)
	rchild := NewNode(3)
	root.left = lchild
	root.right = rchild
	b.root = root

	return b
}

type Btree struct {
	root *Node
}

type Node struct {
	left, right *Node
	data        int
}

type Compare func(i, j int) int8

/*
 Given a binary tree, return true if a node
 with the target data is found in the tree. Recurs
 down the tree, chooses the left or right
 branch by comparing the target to each node.
*/
func (b Btree) Lookup(n *Node, target int) int {
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
			if target < n.data {
				return b.Lookup(n.left, target)
			} else {
				return b.Lookup(n.right, target)
			}
		}
	}
}

/*
 Helper function that allocates a new node
 with the given data and NULL left and right
 pointers.
*/
func NewNode(d int) (n *Node) {
	//n := new(&Node)
	//n.data = d
	//n.left = nil
	//n.right = nil
	//return n
	n.data = d
	return
}

/*
 Give a binary search tree and a number, inserts a new node
 with the given number in the correct place in the tree.
 Returns the new root pointer which the caller should
 then use (the standard trick to avoid using reference
 parameters).
*/
//func Insert(n *Node, d int) Node {
//	// 1. If the tree is empty, return a new, single node
//	if n == nil {
//		return NewNode(d)
//	} else {
//		// 2. Otherwise, recur down the tree
//		if d <= n.data {
//			n.left = Insert(n.left, d)
//		} else {
//			n.right = Insert(n.right, d)
//		}
//		return n
//	}
//}
