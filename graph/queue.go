package graph

import (
	"container/heap"
)

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	//item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Update(i *Item, v Vertex, p int) {
	heap.Remove(pq, i.index)
	i.value = v
	i.priority = p
	heap.Push(pq, i)
}

func (pq PriorityQueue) GetItemForVertex(v Vertex) (i Item) {
	for i := 0; i < len(pq); i++ {
		if pq[i].value == v {
			return *pq[i]
		}
	}
	return
}
