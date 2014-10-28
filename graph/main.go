package main

import (
	"./edge"
	"./graph"
	"./vertex"
	"fmt"
)

func main() {
	fmt.Println("   let's start")

	var g = graph.Graph{}
	//START VERTICES
	var vertex1 = vertex.Vertex{
		Name: "1",
	}
	g.vertices = append(g.vertices, vertex1)

	var vertex2 = vertex.Vertex{
		Name: "2",
	}
	g.vertices = append(g.vertices, vertex2)

	var vertex3 = vertex.Vertex{
		Name: "3",
	}
	g.vertices = append(g.vertices, vertex3)

	var vertex4 = vertex.Vertex{
		Name: "4",
	}
	g.vertices = append(g.vertices, vertex4)

	var vertex5 = vertex.Vertex{
		Name: "5",
	}
	g.vertices = append(g.vertices, vertex5)

	var vertex6 = vertex.Vertex{
		Name: "6",
	}
	g.vertices = append(g.vertices, vertex6)
	//END VERTICES

	//START EDGES
	var edge12 = edge.Edge{
		Weight:  7,
		Vertex1: vertex1,
		Vertex2: vertex2,
	}
	g.edges = append(g.edges, edge12)

	var edge13 = edge.Edge{
		Weight:  9,
		Vertex1: vertex1,
		Vertex2: vertex3,
	}
	g.edges = append(g.edges, edge13)

	var edge16 = edge.Edge{
		Weight:  14,
		Vertex1: vertex1,
		Vertex2: vertex6,
	}
	g.edges = append(g.edges, edge16)

	var edge23 = edge.Edge{
		Weight:  10,
		Vertex1: vertex2,
		Vertex2: vertex3,
	}
	g.edges = append(g.edges, edge23)

	var edge24 = edge.Edge{
		Weight:  15,
		Vertex1: vertex2,
		Vertex2: vertex4,
	}
	g.edges = append(g.edges, edge24)

	var edge34 = edge.Edge{
		Weight:  11,
		Vertex1: vertex3,
		Vertex2: vertex4,
	}
	g.edges = append(g.edges, edge34)

	var edge36 = edge.Edge{
		Weight:  2,
		Vertex1: vertex3,
		Vertex2: vertex6,
	}
	g.edges = append(g.edges, edge36)

	var edge45 = edge.Edge{
		Weight:  6,
		Vertex1: vertex4,
		Vertex2: vertex5,
	}
	g.edges = append(g.edges, edge45)

	var edge56 = edge.Edge{
		Weight:  9,
		Vertex1: vertex5,
		Vertex2: vertex6,
	}
	g.edges = append(g.edges, edge56)

	shortestPaths := graph.Dijkstra(g, vertex1)

	fmt.Println(shortestPaths)

	u := vertex5
	var path []vertex.Vertex
	for {
		currentVertex := shortestPaths[u]
		path = append(path, u)
		if currentVertex.Name == "" {
			break
		}
		u = currentVertex
	}
	fmt.Println(path)
}
