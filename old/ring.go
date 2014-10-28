package main

import (
	"math"
)

type Ring struct {
	start  *Node
	length int
	init   bool
}

func NewRing(n *Node) Ring {
	return Ring{n, 1, true}
}

func (r *Ring) Unset() {
	r.init = false
}

func (r *Ring) MoveAllNodes(s *Sample, gain float64) {
	current := r.start
	best := r.FindMinimum(s)
	for i := 0; i < r.length; i++ {
		current.Move(s, r.F(gain, best.Distance(current, r.length)))
		current = current.right
	}
}

func (r *Ring) F(gain float64, n float64) float64 {
	return (0.70710678 * math.Exp(-(n*n)/(gain*gain)))
}

func (r *Ring) FindMinimum(s *Sample) Node {
	actual := 0.00
	node := r.start
	best := node
	min := node.Potential(s)
	for i := 1; i < r.length; i++ {
		node = node.right
		actual = node.Potential(s)
		if actual < min {
			min = actual
			best = node
		}
	}
	best.isWinner++
	return *best
}

func (r *Ring) DeleteNode(n *Node) {
	previous := n.left
	next := n.right
	if previous.init != false {
		previous.right = next
	}
	if next.init != false {
		next.left = previous
	}
	if next == n {
		next.init = false
	}
	if r.start == n {
		r.start = next
	}
	r.length--
}

func (r *Ring) DuplicateNode(n *Node) {
	a := NewNode(n.x, n.y)
	next := n.left
	next.right = a
	n.left = a
	n.inhibitation = 1
	a.left = next
	a.right = n
	a.inhibitation = 1
	a.init = n.init
	r.length++
}

func (r *Ring) TourLenght() float64 {
	dist := 0.00
	curr := r.start
	prev := curr.left

	for i := 0; i < r.length; i++ {
		x := math.Pow(float64(curr.x-prev.x), 2)
		y := math.Pow(float64(curr.y-prev.y), 2)
		dist += math.Sqrt(x + y)
		curr = prev
		prev = prev.left
	}
	return dist
}
