package slice

import (
	"log"
)

type intSlice []int

func main() {
	log.Print("========> Insertion Sort <========")
	input := intSlice{4, 2, 4, 1, 2, 6, 8, 19, 3}
	log.Print(" input: ", input)
	pop := input.Pop()
	log.Print("output: ", pop)
	log.Print(" input: ", input)
}

// incomplete as it only return the value and doesn't
// touch the slice itself
func (s intSlice) Pop() int {
	//log.Print(" input: ", s)
	//length := len(s)
	//output := s[length-1]
	//s = s[0 : length-2 ]
	//return output
	x := s[len(s)-1]
	s = s[:len(s)-1]
	return x
}
