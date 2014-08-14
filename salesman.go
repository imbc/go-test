// see http://blog.johannes-beck.name/?p=222

package main

import (
	"log"
	"math/rand"
)

func main() {
	a := 47
	b := 6
	c := a % b
	log.Print(a, b, c)
	log.Print("foo")
}

type TravelingSalesman struct {
	N, cycle, maxCycle, update int
	cities                     []Sample
	neurons                    Ring
	alpha, gain, lastLength    float64
	isRunning                  bool
}

func (tsp *TravelingSalesman) CreateFirstNeuron() {
	start := NewNode(0.5, 0.5)
	tsp.neurons = NewRing(start)
}

func (tsp *TravelingSalesman) DeleteAllNeurons() {
	if tsp.neurons.init != false {
		for &tsp.neurons.start != nil {
			tsp.neurons.DeleteNode(tsp.neurons.start)
		}
	}
	tsp.neurons.init = false
}

func (tsp *TravelingSalesman) Print() {
	log.Print("TSP: N = ", tsp.N, ", cycle = ", tsp.cycle, ", lastLength = ", tsp.lastLength)
	for i := 0; i < len(tsp.cities); i++ {
		c := tsp.cities[i]
		log.Print("City: ", i, " (", c.x, ",", c.y, ")")
	}
	n := tsp.neurons.start
	for i := 0; i < tsp.neurons.length; i++ {
		log.Print("Node: ", i, " (", n.x, ",", n.y, ")")
		n = n.right
	}
}

func (tsp *TravelingSalesman) CreateRandomCities() {
	tsp.cities = make([]Sample, tsp.N)
	for i := 0; i < tsp.N; i++ {
		c := NewSample(rand.Float64(), rand.Float64())
		tsp.cities[i] = c
	}
}

func (tsp *TravelingSalesman) Stop() {
	tsp.isRunning = false
	tsp.Repaint()
	//tsp.cities = null
	tsp.DeleteAllNeurons()
}

func (tsp *TravelingSalesman) Init() {
	tsp.cycle = 0
	tsp.lastLength = 0
	tsp.CreateFirstNeuron()
	tsp.CreateRandomCities()
	// call for new Canvas
	tsp.Repaint()
}

func (tsp *TravelingSalesman) Run() {
	if tsp.neurons.init != false {
		if tsp.cycle < tsp.maxCycle && tsp.isRunning == true {
			// setting the interval
		}
		if tsp.isRunning == true {
			tsp.isRunning = false
			tsp.Repaint()
		}
	}
}

func (tsp *TravelingSalesman) Start() {
	tsp.Stop()
	tsp.Init()
	tsp.isRunning = true
	tsp.Run()
}

func (tsp *TravelingSalesman) SurveyRun() bool {
	done := false
	if tsp.neurons.init != false {
		for i := 0; i < len(tsp.cities); i++ {
			tsp.neurons.MoveAllNodes(&tsp.cities[i], tsp.gain)
		}
	}
	tsp.SurveyFinish()
	tsp.gain = tsp.gain * (1 - tsp.alpha)
	tsp.cycle++
	if tsp.cycle%tsp.update == 0 {
		length := tsp.neurons.TourLenght()
		tsp.Repaint()
		if length == tsp.lastLength {
			done = true
		} else {
			tsp.lastLength = length
		}
	}
	return done
}

func (tsp *TravelingSalesman) SurveyFinish() {
	if tsp.neurons.init == false {
		return
	}
	node := tsp.neurons.start
	for i := 0; i < tsp.neurons.length; i++ {
		node.inhibitation = 0
		switch node.isWinner {
		case 0:
			node.life--
			if node.life == 0 {
				tsp.neurons.DeleteNode(node)
			}
		case 1:
			node.life = 3
		default:
			node.life = 3
			tsp.neurons.DuplicateNode(node)
		}
		node.isWinner = 0
		node = node.right
	}
}

func (tsp *TravelingSalesman) Repaint() {

}
