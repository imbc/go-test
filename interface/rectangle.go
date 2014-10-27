/**
 * http://golangtutorials.blogspot.co.uk/2011/06/interfaces-in-go.html
 */
package main

import "fmt"

//Shaper is an interface and has a single function Area that returns an int.
type Shaper interface {
	Area() int
}

//define a Rectangle struct that has a length and a width
type Rectangle struct {
	length, width int
}

//write a function Area that can apply to a Rectangle type
//this function Area works on the type Rectangle and has the same function
//signature defined in the interface Shaper.  Therefore, Rectangle now
//implements the interface Shaper.
func (r Rectangle) Area() int {
	return r.length * r.width
}

type Square struct {
	side int
}

//the Square type also has an Area function and therefore, it too, implements the Shaper interface
func (sq Square) Area() int {
	return sq.side * sq.side
}

func main() {
	r := Rectangle{length: 5, width: 3} //define a new Rectangle instance with values for its properties
	fmt.Println("Rectangle details are: ", r)
	fmt.Println("Rectangle's area is: ", r.Area())
	fmt.Println()

	var s Shaper
	s = r //equivalent to "s = Shaper(r)" since Go identifies r matches the Shaper interface
	fmt.Println("Area of Shaper(Rectangle): ", s.Area())
	fmt.Println()

	q := Square{side: 5}
	fmt.Println("Square details are: ", q)
	s = q //equivalent to "s = Shaper(q)
	fmt.Println("Area of Shaper(Square): ", s.Area())
	fmt.Println()
	shapesArr := [...]Shaper{r, q}

	fmt.Println("Looping through shapes for area ...")
	for n, _ := range shapesArr {
		fmt.Println("Shape details: ", shapesArr[n])
		fmt.Println("Area of this shape is: ", shapesArr[n].Area())
	}
}
