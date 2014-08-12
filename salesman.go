// see http://blog.johannes-beck.name/?p=222

package main

import (
	"log"
)

func main() {
	log.Print("foo")
}

type TravelingSalesman struct {
	N, cycle, maxCycle, update, lastLength int
	cities                                 []Node
	neurons                                Ring
	alpha, gain                            float64
	isRunning                              bool
}

func (tsp TravelingSalesman) CreateFirstNeuron() {
	start := NewNode(0.5, 0.5)
	tsp.neurons = NewRing(start)
}

func (tsp TravelingSalesman) DeleteAllNeurons() {
	if tsp.neurons.init != false {
		for &tsp.neurons.start != nil {
			tsp.neurons.DeleteNode(tsp.neurons.start)
		}
	}
	tsp.neurons.init = false
}

func (tsp TravelingSalesman) Print() {
	log.Print("TSP: N = ", tsp.N, ", cycle = ", tsp.cycle, ", lastLength = ", tsp.lastLength)
	for i := 0; i < len(tsp.cities); i++ {
		c := tsp.cities[i]
		log.Print("City: ", i, " (", c.x, ",", c.y, ")")
	}
	//n := tsp.neurons.start
}
