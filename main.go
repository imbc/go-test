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
	width, height int              = 640, 640
	colors        []colorful.Color = colorful.FastHappyPalette(7)

	max_x = 1024
	max_y = 1024
)

func main() {
	i := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(i, i.Bounds(), &image.Uniform{black}, image.ZP, draw.Src)
	gc := draw2d.NewGraphicContext(i)
	gc.MoveTo(10.0, 10.0)
	gc.LineTo(100.0, 10.0)
	gc.Stroke()
	//width, height := 640, 640
	//canvas := NewCanvas(image.Rect(0, 0, width, height))
	//canvas.DrawRect(
	//	color.RGBA{0, 0, 0, 255},
	//	Point{0, 0},
	//	Point{float64(width), float64(height)})
	rand.Seed(time.Now().UTC().UnixNano())

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
