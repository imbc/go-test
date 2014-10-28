package edge

import (
	"../vertex"
)

type Edge struct {
	weight  int
	vertex1 vertex.Vertex
	vertex2 vertex.Vertex
}

/**
 * check if the edge is connected to the given vertex
 */
func (e *Edge) Connected(v vertex.Vertex) bool {
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
func (e *Edge) OtherSide(v vertex.Vertex) (r vertex.Vertex) {
	if e.vertex1 == v {
		return e.vertex2
	}
	if e.vertex2 == v {
		return e.vertex1
	}
	return
}
