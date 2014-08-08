/**
 * Sorting algorithm based on the version from
 * http://www.stoimen.com/blog/
 *
 * Nothing very original, but at least it is both
 * a refresh for the theory and a learning to for
 * Golang.
 */

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
	quick := QuickSort(input)
	log.Print("quicksort array: ", quick)
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
	output := make([]int, 0)
	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			output = append(output, left[0])
			left = left[1:]
		} else {
			output = append(output, right[0])
			right = right[1:]
		}
	}
	output = append(output, left...)
	output = append(output, right...)
	return output
}

func QuickSort(slice []int) []int {
	if len(slice) == 0 {
		return slice
	}
	pivot := slice[0]
	left, right := make([]int, 0), make([]int, 0)
	for i := 1; i < len(slice); i++ {
		if slice[i] < pivot {
			left = append(left, slice[i])
		} else {
			right = append(right, slice[i])
		}
	}
	output := append(QuickSort(left), pivot)
	output = append(output, QuickSort(right)...)
	return output
}
