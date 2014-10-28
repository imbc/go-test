/**
 * Sorting algorithm based on the version from
 * http://www.stoimen.com/blog/
 *
 * Nothing very original, but at least it is both
 * a refresh for the theory and a learning to for
 * Golang.
 */

package sort

import (
	"log"
	"math"
)

func Print() {
	log.Print("========> Insertion Sort <========")
	iInsertion := []int{4, 2, 4, 1, 2, 6, 8, 19, 3}
	log.Print(" input: ", iInsertion)
	insertion := InsertionSort(iInsertion)
	log.Print("output: ", insertion)

	log.Print("========> Merge Sort <========")
	iMerge := []int{6, 5, 3, 1, 8, 7, 2, 4}
	log.Print(" input: ", iMerge)
	merge := MergeSort(iMerge)
	log.Print("output: ", merge)

	log.Print("========> Quick Sort <========")
	iQuick := []int{5, 3, 9, 8, 7, 2, 4, 1, 6, 5}
	log.Print(" input: ", iQuick)
	quick := QuickSort(iQuick)
	log.Print("output: ", quick)

	log.Print("========> Bubble Sort <========")
	iBubble := []int{6, 5, 3, 1, 8, 7, 2, 4}
	log.Print(" input: ", iBubble)
	bubble := QuickSort(iBubble)
	log.Print("output: ", bubble)

	log.Print("========> Radix Sort <========")
	iRadix := []int{4, 3, 5, 9, 7, 2, 4, 1, 6, 5}
	log.Print(" input: ", iRadix)
	radix := RadixSort(iRadix)
	log.Print("output: ", radix)

	log.Print("========> Shell Sort <========")
	iShell := []int{6, 5, 3, 1, 8, 7, 2, 4}
	log.Print(" input: ", iShell)
	shell := ShellSort(iShell)
	log.Print("output: ", shell)

	log.Print("========> Bucket Sort <========")
	iBucket := []int{4, 3, 5, 9, 7, 8, 7, 2, 4}
	log.Print(" input: ", iBucket)
	bucket := BucketSort(iBucket)
	log.Print("output: ", bucket)
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

func ShellSort(slice []int) []int {
	gaps := []int{1, 2, 3, 4, 6}
	gap := gaps[len(gaps)-1]     // pop
	gaps = gaps[0 : len(gaps)-1] // pop
	temp := 0
	output := make([]int, len(slice))
	for gap > 0 {
		for i := gap; i < len(slice); i++ {
			temp = slice[i]
			j := i
			for j >= gap && output[j-gap] > temp {
				output[j] = output[j-gap]
				j -= gap
			}
			output[j] = temp
		}
		if len(gaps) > 0 {
			gap = gaps[len(gaps)-1]
			gaps = gaps[0 : len(gaps)-1]
		} else {
			gap = 0
		}
	}
	return output
}

// incomplete as it only return the value and doesn't touch the slice itself
func Pop(slice []int) int {
	output := slice[len(slice)-1]
	slice = slice[0 : len(slice)-2]
	return output
}

func Reverse(slice []int) []int {
	output := make([]int, 0)
	for i := len(slice) - 1; i > -1; i-- {
		output = append(output, slice[i])
	}
	return output
}

type Bucket struct {
	bit      []int
	min, max int
}

// seems to be a good candidate for goroutine
func BucketSort(slice []int) []int {
	n := math.Floor(math.Sqrt(float64(len(slice))))
	div := math.Floor(math.Sqrt(float64(Sum(slice)) / float64(n)))
	min := Min(slice)
	buckets := make([]Bucket, 0)
	//lets make & populate the buckets
	for i := 1; i < int(n)+1; i++ {
		max := int(math.Floor(div * float64(i)))
		b := Bucket{make([]int, 0), min, max}
		for _, val := range slice {
			if min <= val && val <= max {
				b.bit = append(b.bit, val)
			}
		}
		buckets = append(buckets, b)
		min = max + 1
	}

	output := make([]int, 0)
	for i := 0; i < len(buckets); i++ {
		if len(buckets) > 0 {
			output = append(output, InsertionSort(buckets[i].bit)...)
		}
	}
	return output
}

func Sum(slice []int) int {
	output := 0
	for _, val := range slice {
		output += val
	}
	return output
}
