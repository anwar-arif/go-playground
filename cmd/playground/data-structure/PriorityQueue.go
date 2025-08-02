package data_structure

import (
	"container/heap"
	"fmt"
)

// MinPriorityQueue is a min-heap for int64 values
type MinPriorityQueue struct {
	items []int64
}

// Implement heap.Interface
func (pq *MinPriorityQueue) Len() int           { return len(pq.items) }
func (pq *MinPriorityQueue) Less(i, j int) bool { return pq.items[i] < pq.items[j] }
func (pq *MinPriorityQueue) Swap(i, j int)      { pq.items[i], pq.items[j] = pq.items[j], pq.items[i] }

func (pq *MinPriorityQueue) Push(x interface{}) {
	pq.items = append(pq.items, x.(int64))
}

func (pq *MinPriorityQueue) Pop() interface{} {
	old := pq.items
	n := len(old)
	item := old[n-1]
	pq.items = old[0 : n-1]
	return item
}

// NewMinPQ creates a new min priority queue
func NewMinPQ() *MinPriorityQueue {
	pq := &MinPriorityQueue{items: make([]int64, 0)}
	heap.Init(pq)
	return pq
}

// PushValue adds a value to the priority queue
func (pq *MinPriorityQueue) PushValue(value int64) {
	heap.Push(pq, value)
}

// PopValue removes and returns the smallest value
func (pq *MinPriorityQueue) PopValue() (int64, bool) {
	if pq.IsEmpty() {
		return 0, false
	}
	return heap.Pop(pq).(int64), true
}

// Peek returns the smallest value without removing it
func (pq *MinPriorityQueue) Peek() (int64, bool) {
	if pq.IsEmpty() {
		return 0, false
	}
	return pq.items[0], true
}

// IsEmpty returns true if the queue is empty
func (pq *MinPriorityQueue) IsEmpty() bool {
	return len(pq.items) == 0
}

// Size returns the number of elements
func (pq *MinPriorityQueue) Size() int {
	return len(pq.items)
}

func RunPriorityQueue() {
	fmt.Println("=== Testing Priority Queue Correctness ===\n")

	// Test the exact case you mentioned
	fmt.Println("Test: Push 0-9, then pop all (Min-Heap):")
	minPQ := NewMinPQ()

	// Push 0 to 9
	fmt.Print("Pushing: ")
	for i := int64(0); i < 10; i++ {
		minPQ.PushValue(i)
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// Pop all values
	fmt.Print("Popping: ")
	result := make([]int64, 0)
	for !minPQ.IsEmpty() {
		if val, ok := minPQ.PopValue(); ok {
			result = append(result, val)
			fmt.Printf("%d ", val)
		}
	}
	fmt.Printf("\nResult: %v\n", result)

	// Verify it's sorted
	isCorrect := true
	for i := 0; i < len(result); i++ {
		if result[i] != int64(i) {
			isCorrect = false
			break
		}
	}
	fmt.Printf("Correct order: %t\n\n", isCorrect)
}
