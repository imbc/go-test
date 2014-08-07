package main

import (
	"log"
)

func main() {
	input := []int{1, 3, 7, 2, 9, 8, 4, 6, 5, 0}
	log.Print("input array: ", input)
	//log.Print(Shift(input))
	//for key, value := range Shift(input) {
	//	log.Print(key, value)
	//}
	//insertion := InsertionSort(input)
	//log.Print(insertion)
	//MergeSort(input)
	merge := MergeSort(input)

	log.Print("merged array: ", merge)
}

func InsertionSort(slice []int) []int {
	for i := 1; i < len(slice); i++ {
		elt := slice[i]
		pos := i
		for pos >= 1 && slice[pos-1] > elt {
			slice[pos] = slice[pos-1]
			pos--
		}
		slice[pos] = elt
	}
	return slice
}

func MergeSort(slice []int) []int {
	if len(slice) <= 1 {
		return slice
	}
	// split the slice in 2
	left := slice[0 : len(slice)/2]
	right := slice[len(slice)/2:]
	//log.Print("left + right ", left, right)
	// repeat
	left = MergeSort(left)
	right = MergeSort(right)
	// merge the slice
	//log.Print("left + right ", left, right)
	output := Merge(left, right)
	return output
}

func Merge(left []int, right []int) []int {

	length := 0
	if len(left) > len(right) {
		length = len(left)
	} else {
		length = len(right)
	}
	//log.Print("max length: ", length)
	result := make([]int, length)
	//log.Print(result)
	for len(left) > 0 && len(right) > 0 {

		//log.Print("left array: ", left)
		//log.Print("right array: ", right)

		// get the head of each array
		// compare them
		// taken the smaler item
		// add it to the result array
		// shift the array of origin by one
		//log.Print(i)
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
		//log.Print(result)
	}
	result = append(result, left...)
	result = append(result, right...)
	//log.Print(result)
	return result
}

func Shift(slice []int) []int {
	slice = slice[1:]
	return slice
}
