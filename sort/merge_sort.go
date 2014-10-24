package sort

const TEMP_SIZE = 1024
const INSERTION_SORT_SIZE = 16

type array []*int
type obj struct {
	value    int
	identity string
}

type comparator interface {
	objects(a obj, b obj) obj
}

func (c comparator) Compare(a obj, b obj) bool {
	if a > b {
		return true
	} else {
		return false
	}
}

func (c comparator) Equals(a obj, b obj) bool {
	if a == b {
		return true
	} else {
		return false
	}
}

func (a array) SwapBlocks(from, second, to int) {
	len1 := second - from
	len2 := to - second + 1
	if len1 == 0 || len2 == 0 {
		return
	}
	// ...
}

func (a array) SortArray(c comparator) {

}

func (a array) ReverseBlock(from, to int) {
	for from < to {
		old := a[from]
		//a[from++] = a[to]
		//a[to--] = old
	}
}

func (a array) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
