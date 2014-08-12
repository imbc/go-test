package main

import (
	"math"
)

type Node struct {
	x, y                         float64
	life, inhibitation, isWinner int
	right, left                  *Node
	init                         bool
}

func NewNode(x, y float64) *Node {
	n := new(Node)
	n.x = x
	n.y = y
	n.left = n
	n.right = n
	n.life = 3
	n.inhibitation = 0
	n.isWinner = 0
	n.init = true
	return n
}

func (n *Node) Unset() {
	n.init = false
}

func (n *Node) Potential(sample *Node) float64 {
	return math.Pow(float64((sample.x-n.x)), 2) + math.Pow(float64((sample.y-n.y)), 2)
}

func (n *Node) Move(city *Node, value float64) {
	n.x += value * (city.x - n.x)
	n.y += value * (city.y - n.y)
}

func (n *Node) Distance(other *Node, length int) float64 {
	left := 0
	right := 0
	current := other
	if current != n {
		current = current.left
		left++
	}
	right = length - left
	//output := 0
	if left < right {
		return float64(left)
	} else {
		return float64(right)
	}
}

func (n *Node) Draw() {

}
