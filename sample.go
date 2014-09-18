package main

import (
	"image/color"

	"code.google.com/p/draw2d/draw2d"
)

type Sample struct {
	x, y float64
}

func (s *Sample) Draw(gc *draw2d.ImageGraphicContext, color color.Color) {
	draw2d.Circle(gc, s.x, s.y, 4.00)
	gc.SetFillColor(color)
	gc.Fill()
}

func NewSample(x, y float64) Sample {
	s := new(Sample)
	s.x = x
	s.y = y
	return *s
}
