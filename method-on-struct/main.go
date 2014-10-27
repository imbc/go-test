package main

import (
	"fmt"
	"time"
)

type Rectangle struct {
	length, width int
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

func (r Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}

func (r Rectangle) Area_by_value() int {
	return r.length * r.width
}

func (r *Rectangle) Area_by_reference() int {
	return r.length * r.width
}

type myTime struct {
	time.Time //anonymous field
}

func (t myTime) first5Chars() string {
	return t.Time.String()[0:5]
}

func main() {
	r1 := Rectangle{4, 3}
	fmt.Println("Rectangle is: ", r1)
	fmt.Println("Rectangle area is: ", r1.Area())
	fmt.Println("Rectangle perimeter is: ", r1.Perimeter())
	fmt.Println("Rectangle area is: ", r1.Area_by_value())
	fmt.Println("Rectangle area is: ", r1.Area_by_reference())
	fmt.Println("Rectangle area is: ", (&r1).Area_by_value())
	fmt.Println("Rectangle area is: ", (&r1).Area_by_reference())
	//time.LoadLocation("UTC")
	m := myTime{*time.Local()}                     //since time.LocalTime returns an address, we convert it to a value with *
	fmt.Println("Full time now:", m.String())      //calling existing String method on anonymous Time field
	fmt.Println("First 5 chars:", m.first5Chars()) //calling myTime.first5Chars
}
