package graph

import (
	"fmt"
)

type Vertex struct {
	name string
}

func (v *Vertex) Print() {
	fmt.Println("> Vertex: ", v.name)
}
