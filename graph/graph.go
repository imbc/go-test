package graph

import (
	"container/heap"
)

type Graph struct {
	vertices []Vertex
	edges    []Edge
}

func (g *Graph) AddVertex(n string) (v Vertex) {
	v.name = n
	g.vertices = append(g.vertices, v)
	return v
}

func (g *Graph) AddEdge(w int, v1 Vertex, v2 Vertex) (e Edge) {
	e.weight = w
	e.vertex1 = v1
	e.vertex2 = v2
	g.edges = append(g.edges, e)
	return e
}

func (g *Graph) Neighbours(v Vertex) (r []Vertex) {
	for _, edge := range g.edges {
		if edge.Connected(v) {
			var neighbour Vertex = edge.Opposite(v)
			r = append(r, neighbour)
		}
	}
	return r
}

func (g *Graph) Edges(origin Vertex, dest Vertex) (e Edge) {
	for _, edge := range g.edges {
		if edge.Connected(origin) {
			if edge.vertex1 == dest {
				return edge
			}
			if edge.vertex2 == dest {
				return edge
			}
		}
	}
	return
}

func Dijkstra(g Graph, source Vertex) (r map[Vertex]Vertex) {
	pq := &PriorityQueue{}
	heap.Init(pq)

	distance := make(map[Vertex]int)
	previous := make(map[Vertex]Vertex)

	distance[source] = 0

	for _, v := range g.vertices {
		if v != source {
			distance[v] = 9999
		}
		item := &Item{
			value:    v,
			priority: distance[v],
		}
		heap.Push(pq, item)
	}

	for pq.Len() > 0 {
		currentItem := heap.Pop(pq).(*Item)
		currentVertex := currentItem.value

		for _, neighbour := range g.Neighbours(currentVertex) {
			var e = g.Edges(currentVertex, neighbour)

			if &e != nil {
				var alt int = distance[currentVertex] + e.weight
				if alt < distance[neighbour] {
					distance[neighbour] = alt
					previous[neighbour] = currentVertex
					neighbourItem := pq.GetItemForVertex(neighbour)
					pq.Update(&neighbourItem, neighbour, alt)
				}
			}
		}
	}
	return previous
}
