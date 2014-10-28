package graph

import (
	"fmt"
)

type Edge struct {
	weight  int
	vertex1 Vertex
	vertex2 Vertex
}

/**
 * check if the edge is connected to the given vertex
 */
func (e *Edge) Connected(v Vertex) bool {
	if e.vertex1 == v {
		return true
	}
	if e.vertex2 == v {
		return true
	}
	return false
}

/**
 * return the opposite vertex
 */
func (e *Edge) Opposite(v Vertex) (r Vertex) {
	if e.vertex1 == v {
		return e.vertex2
	}
	if e.vertex2 == v {
		return e.vertex1
	}
	return
}

func (e *Edge) Print() {
	fmt.Println("Edge connecting: ", e.vertex1.name, " and ", e.vertex2.name)
	e.vertex1.Print()
	e.vertex2.Print()
}
