package graph

import (
	"../edge"
	"../queue"
	"../vertex"
	"container/heap"
)

type Graph struct {
	vertices []vertex.Vertex
	edges    []edge.Edge
}

func (g *Graph) Neighbours(v vertex.Vertex) (r []vertex.Vertex) {
	//var neighbours []vertex.Vertex
	for _, edge := range g.edges {
		if edge.Connected(v) {
			var neighbour vertex.Vertex = edge.OtherSide(v)
			r = append(r, neighbour)
		}
	}
	return r
}

func (g *Graph) Edges(origin vertex.Vertex, dest vertex.Vertex) (r edge.Edge) {
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

func Dijkstra(g Graph, source vertex.Vertex) (r map[vertex.Vertex]vertex.Vertex) {
	pq := &queue.PriorityQueue{}
	heap.Init(pq)

	distance := make(map[vertex.Vertex]int)
	previous := make(map[vertex.Vertex]vertex.Vertex)

	distance[source] = 0

	for _, v := range g.vertices {
		if v != source {
			distance[v] = 9999
		}
		item := &queue.Item{
			Value:    v,
			Priority: distance[v],
		}
		heap.Push(pq, item)
	}

	for pq.Len() > 0 {
		currentItem := heap.Pop(pq).(*queue.Item)
		currentVertex := currentItem.Value

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
