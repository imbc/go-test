package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"math/rand"
	"os"
	//"sort"
	"time"

	"code.google.com/p/draw2d/draw2d"
	"github.com/demdxx/go-colorful"
)

/** Sets Colors for easy use **/
var (
	white         color.Color      = color.RGBA{255, 255, 255, 255}
	black         color.Color      = color.RGBA{0, 0, 0, 255}
	teal          color.Color      = color.RGBA{0, 200, 200, 255}
	red           color.Color      = color.RGBA{200, 30, 30, 255}
	green         color.Color      = color.RGBA{0, 255, 0, 255}
	blue          color.Color      = color.RGBA{0, 0, 255, 255}
	orange        color.Color      = color.RGBA{255, 168, 0, 255}
	lred          color.Color      = color.RGBA{200, 30, 30, 0x80}
	lgreen        color.Color      = color.RGBA{0, 255, 0, 0x80}
	lblue         color.Color      = color.RGBA{0, 0, 255, 0x80}
	lorange       color.Color      = color.RGBA{255, 168, 0, 0x80}
	width, height float64          = 640.00, 640.00
	colors        []colorful.Color = colorful.FastHappyPalette(7)

	max_x = 1024
	max_y = 1024
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	p := new(Point)
	p.x = x
	p.y = y
	return p
}

func (p *Point) Draw(gc *draw2d.ImageGraphicContext, color color.Color) {
	draw2d.Circle(gc, p.x, p.y, 4.00)
	gc.SetFillColor(color)
	gc.Fill()
}

type Region struct {
	bound_a, bound_b Point
	pop              []Point
}

func NewRegion(p1, p2 *Point) *Region {
	r := new(Region)
	r.bound_a = *p1
	r.bound_b = *p2
	return r
}

func (r *Region) Draw(gc *draw2d.ImageGraphicContext, color color.Color) {
	draw2d.Rect(gc, r.bound_a.x, r.bound_a.y, r.bound_b.x, r.bound_b.y)
	gc.SetFillColor(color)
	gc.Fill()
}

func (r *Region) DrawPop(gc *draw2d.ImageGraphicContext, color color.Color) {
	for _, pt := range r.pop {
		pt.Draw(gc, color)
	}
}

func Distance(x, y float64) float64 {
	if x > y {
		return x - y
	} else {
		return y - x
	}
}

func (r Region) Contains(p *Point) bool {
	dx := Distance(r.bound_a.x, r.bound_b.x)
	cx1 := Distance(p.x, r.bound_a.x)
	cx2 := Distance(p.x, r.bound_b.x)
	dy := Distance(r.bound_a.y, r.bound_b.y)
	cy1 := Distance(p.y, r.bound_a.y)
	cy2 := Distance(p.y, r.bound_b.y)
	if dx == cx1+cx2 && dy == cy1+cy2 {
		return true
	} else {
		return false
	}
}

func Random(min, max int) int {
	return rand.Intn(max-min) + min
}

func (r *Region) AddPoint(pt Point) {
	r.pop = append(r.pop, pt)
}

func (r *Region) Populate(nb int) {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < nb; i++ {
		a := float64(Random(int(r.bound_a.x), int(r.bound_b.x)))
		b := float64(Random(int(r.bound_a.y), int(r.bound_b.y)))
		r.AddPoint(*NewPoint(a, b))
	}
}

func main() {
	i := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	draw.Draw(i, i.Bounds(), &image.Uniform{black}, image.ZP, draw.Src)
	gc := draw2d.NewGraphicContext(i)
	gc.MoveTo(10.0, 10.0)
	gc.LineTo(100.0, 10.0)
	gc.Stroke()

	a := NewPoint(0.00, 0.00)
	b := NewPoint(width, height)
	c := NewPoint(width/2, 0.00)
	d := NewPoint(width, height/2)
	e := NewPoint(0.00, height/2)
	f := NewPoint(width/2, height)
	g := NewPoint(width/2, height/2)

	r1 := NewRegion(a, g)
	r2 := NewRegion(c, d)
	r3 := NewRegion(e, f)
	r4 := NewRegion(g, b)

	r1.Populate(12)
	r2.Populate(12)
	r3.Populate(12)
	r4.Populate(12)

	r1.DrawPop(gc, lred)
	r2.DrawPop(gc, lblue)
	r3.DrawPop(gc, lgreen)
	r4.DrawPop(gc, lorange)

	//x1 := NewPoint(340, 560)
	//log.Print(r1.Contains(x1))
	//log.Print(r2.Contains(x1))
	//log.Print(r3.Contains(x1))
	//log.Print(r4.Contains(x1))
	//width, height := 640, 640
	//canvas := NewCanvas(image.Rect(0, 0, width, height))
	//canvas.DrawRect(
	//	color.RGBA{0, 0, 0, 255},
	//	Point{0, 0},
	//	Point{float64(width), float64(height)})
	//rand.Seed(time.Now().UTC().UnixNano())

	/*
	 * Create and populate the slice of Nodes
	 */
	//n := 49
	//peers := 7
	//nodes := make([]*Node, n)
	//group := 0
	//groups := make([]int, peers)
	//for i := 0; i < n; i++ {
	//	node := NewNode(peers, canvas)
	//	if i%peers == 0 {
	//		group++
	//	}
	//	node.Group = group
	//	nodes[i] = node
	//}

	//Randomly point Nodes at each other
	//for _, node := range nodes {
	//	for _, j := range rand.Perm(n)[:peers] {
	//		node.Peers = append(node.Peers, nodes[j])
	//	}
	//}

	// Calculate nearest peers for each node
	// This is pretty ineffecient for large n
	//nodesCopy := make([]*Node, n)
	//copy(nodesCopy, nodes)
	//log.Print("Sorting nodes...")
	//for _, node := range nodes {
	//	// Sort the nodes by distance
	//	sorter := NodeSorter{nodesCopy, node}
	//	sort.Sort(sorter)
	//	node.Peers = append(node.Peers, nodesCopy[1:peers+1]...)
	//}
	log.Print("Nodes sorted")

	// Draw on circles representing nodes
	log.Print("Drawing Nodes")
	//colors := colorful.FastHappyPalette(7)
	//for _, node := range nodes {
	//	c := colors[node.Group-1]
	//	canvas.DrawCircle(
	//		color.RGBA{
	//			uint8(c.R * 255),
	//			uint8(c.G * 255),
	//			uint8(c.B * 255),
	//			255},
	//		node.Position,
	//		5)

	//}
	//canvas.Blur(3, new(WeightFunctionDist))

	// Draw connections between nodes
	//log.Print("Drawing connection")
	//for _, node := range nodes {
	//	for _, peer := range node.Peers[:3] {
	//		canvas.DrawLine(color.RGBA{0, 0, 0, 10}, node.Position, peer.Position)
	//	}
	//}

	// Start sending messages between nodes
	//log.Print("sending message")
	//for i := 0; i < 1; i++ {
	//	nodes[i].Power = 255
	//	go nodes[i].Send()
	//}
	//time.Sleep(time.Second)
	saveToPngFile("img/TestPath.png", i)
}

func saveToPngFile(filePath string, m image.Image) {
	f, err := os.Create(filePath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	b := bufio.NewWriter(f)
	err = png.Encode(b, m)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = b.Flush()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Wrote %s OK.\n", filePath)
}
