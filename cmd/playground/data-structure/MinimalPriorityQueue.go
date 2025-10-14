package data_structure

import (
	"container/heap"
	"fmt"
)

type Item struct {
	value int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].value < pq[j].value
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, &Item{x.(int)})
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

func RunMinimalPriorityQueue() {
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, 1)
	heap.Push(pq, 6)
	heap.Push(pq, 3)

	item := heap.Pop(pq).(*Item)
	fmt.Println(item.value)
}
