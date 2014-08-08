package main

import (
	"log"
)

func main() {
	input := []int{1, 3, 7, 2, 9, 8, 4, 6, 5, 0}
	log.Print("input array: ", input)
	insertion := InsertionSort(input)
	log.Print("insertion array: ", insertion)
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
	// repeat
	left = MergeSort(left)
	right = MergeSort(right)
	// merge the slice
	output := Merge(left, right)
	return output
}

func Merge(left []int, right []int) []int {
	result := make([]int, 0)
	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	result = append(result, left...)
	result = append(result, right...)
	return result
}
