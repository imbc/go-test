package main

type Point struct {
	X, Y float64
}

type Vector struct {
	direction, length int
}

type Matrix struct {
	columns []Column
}

type Column struct {
	element   []int
	dimension int
}

func (m *Matrix) Add(m2 Matrix) {
	for i, e := range m {

	}
}
