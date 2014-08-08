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
	iInsertion := []int{4, 2, 4, 1, 2, 6, 8, 19, 3}
	log.Print("insertion array: ", iInsertion)
	insertion := InsertionSort(iInsertion)
	log.Print(" insertion sort: ", insertion)

	iMerge := []int{6, 5, 3, 1, 8, 7, 2, 4}
	log.Print("    merge array: ", iMerge)
	merge := MergeSort(iMerge)
	log.Print("	 merge sort: ", merge)

	iQuick := []int{5, 3, 9, 8, 7, 2, 4, 1, 6, 5}
	log.Print("    quick array: ", iQuick)
	quick := QuickSort(iQuick)
	log.Print("     quick sort: ", quick)

	iBubble := []int{6, 5, 3, 1, 8, 7, 2, 4}
	log.Print("   bubble array: ", iBubble)
	bubble := QuickSort(iBubble)
	log.Print("    bubble sort: ", bubble)

	iRadix := []int{4, 3, 5, 9, 7, 2, 4, 1, 6, 5}
	log.Print("    radix array: ", iRadix)
	radix := RadixSort(iRadix)
	log.Print("     radix sort: ", radix)

	//iShell := []int{6, 5, 3, 1, 8, 7, 2, 4}
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
	left := slice[0 : len(slice)/2]
	right := slice[len(slice)/2:]
	left = MergeSort(left)
	right = MergeSort(right)
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

func BubbleSort(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		for j := len(slice) - 1; j > i; j-- {
			if slice[j] < slice[j-1] {
				slice[j-1], slice[j] = slice[j], slice[j-1]
			}
		}
	}
	return slice
}

func RadixSort(slice []int) []int {
	max := Max(slice) + 1
	if Min(slice) < 0 {
		max -= Min(slice)
	}
	temp := make([]int, max)
	for _, val := range slice {
		temp[val]++
	}
	output := make([]int, 0)
	for key, val := range temp {
		for ; val > 0; val-- {
			output = append(output, key)
		}
	}
	return output
}

func Max(slice []int) int {
	output := slice[0]
	for _, val := range slice {
		if val >= output {
			output = val
		}
	}
	return output
}

func Min(slice []int) int {
	output := slice[0]
	for _, val := range slice {
		if val <= output {
			output = val
		}
	}
	return output
}
