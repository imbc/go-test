package main

import "fmt"

func main() {
    intarr := [5]int{12, 34, 55, 66, 43}
    slice := intarr[:]
	fmt.Printf("address of first element: %p\n", &slice[0])
	fmt.Printf("address of slice itself:  %p\n", &slice)
    fmt.Printf("the len is %d and cap is %d \n", len(slice), cap(slice))
    fmt.Printf("address of slice %p add of Arr %p\n", &slice, &intarr)
}
