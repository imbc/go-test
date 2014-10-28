package main

import (
	"image"
	"math"
)

type Point struct {
	X, Y float64
}

func (p *Point) Rotate(angle float64) {
	cos, sin := math.Cos(angle), math.Sin(angle)
	p.X, p.Y = p.X*cos+p.Y*sin, p.Y*cos-p.X*sin
}

func (p *Point) Scale(k float64) {
	p.X, p.Y = p.X*k, p.Y*k
}

func (p1 *Point) Add(p2 Point) Point {
	return Point{p1.X + p2.X, p1.Y + p2.Y}
}

func (p1 *Point) Sub(p2 Point) Point {
	return Point{p1.X - p2.X, p1.Y - p2.Y}
}

func (p *Point) Length() float64 {
	return math.Hypot(p.X, p.Y)
}

func (p *Point) toPoint() image.Point {
	return image.Point{int(p.X), int(p.Y)}
}
