package data_structure

import (
	"container/heap"
	"fmt"
)

type PriorityQueue []int64

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i] < pq[j] }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq PriorityQueue) IsEmpty() bool      { return len(pq) == 0 }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(int64))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	item := old[0]
	*pq = old[1:]
	return item
}

func NewPriorityQueue() *PriorityQueue {
	pq := &PriorityQueue{}
	heap.Init(pq)
	return pq
}

func (pq *PriorityQueue) PushVal(val int64) {
	heap.Push(pq, val)
}

func (pq *PriorityQueue) PopVal() (int64, bool) {
	if pq.IsEmpty() {
		return 0, false
	}
	return heap.Pop(pq).(int64), true
}

func RunPriorityQueue() {
	pq := NewPriorityQueue()
	for i := 0; i < 10; i++ {
		pq.Push(int64(i))
	}

	for !pq.IsEmpty() {
		top, _ := pq.PopVal()
		fmt.Println(top)
	}
}
