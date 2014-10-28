package queue

import (
	"../vertex"
	"container/heap"
)

type Item struct {
	Value    vertex.Vertex
	Priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
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

func (pq *PriorityQueue) Update(item *Item, Value vertex.Vertex, Priority int) {
	heap.Remove(pq, item.index)
	item.Value = Value
	item.Priority = Priority
	heap.Push(pq, item)
}

func (pq PriorityQueue) GetItemForVertex(aVertex vertex.Vertex) (item Item) {
	for i := 0; i < len(pq); i++ {
		if pq[i].Value == aVertex {
			return *pq[i]
		}
	}
	return
}
