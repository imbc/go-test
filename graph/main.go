package main

import (
	"fmt"
)

func main() {
	fmt.Println("   let's start")
	fmt.Println(" == VERTICES ==")

	g := Graph{}

	fmt.Println(" == VERTICES ==")
	//START VERTICES
	v1 := g.AddVertex("1")
	v1.Print()
	v2 := g.AddVertex("2")
	v2.Print()
	v3 := g.AddVertex("3")
	v3.Print()
	v4 := g.AddVertex("4")
	v4.Print()
	v5 := g.AddVertex("5")
	v5.Print()
	v6 := g.AddVertex("6")
	v6.Print()
	//END VERTICES

	//START EDGES
	g.AddEdge(7, v1, v2)
	g.AddEdge(9, v1, v3)
	g.AddEdge(14, v1, v6)
	g.AddEdge(10, v2, v3)
	g.AddEdge(15, v2, v4)
	g.AddEdge(11, v3, v4)
	g.AddEdge(2, v3, v6)
	g.AddEdge(6, v4, v5)
	g.AddEdge(9, v5, v6)

	shortestPaths := Dijkstra(g, v1)

	fmt.Println(shortestPaths)

	u := v5
	var path []Vertex
	for {
		currentVertex := shortestPaths[u]
		path = append(path, u)
		if currentVertex.name == "" {
			break
		}
		u = currentVertex
	}
	fmt.Println(path)
}
