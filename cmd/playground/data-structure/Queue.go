package data_structure

import (
	"errors"
	"fmt"
	"sync"
)

// BasicQueue - Simple queue implementation using a slice
type BasicQueue struct {
	items []int
}

func (q *BasicQueue) Enqueue(item int) {
	q.items = append(q.items, item)
}

func (q *BasicQueue) Dequeue() (int, error) {
	if len(q.items) == 0 {
		return 0, errors.New("queue is empty")
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}

func (q *BasicQueue) Front() (int, error) {
	if len(q.items) == 0 {
		return 0, errors.New("queue is empty")
	}
	return q.items[0], nil
}

func (q *BasicQueue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *BasicQueue) Size() int {
	return len(q.items)
}

// ConcurrentQueue - Thread-safe queue implementation
type ConcurrentQueue struct {
	items []int
	mu    sync.Mutex
}

func (q *ConcurrentQueue) Enqueue(item int) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.items = append(q.items, item)
}

func (q *ConcurrentQueue) Dequeue() (int, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.items) == 0 {
		return 0, errors.New("queue is empty")
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}

// Generic Queue - More flexible implementation
type GenericQueue[T any] struct {
	items []T
	mu    sync.RWMutex
}

func (q *GenericQueue[T]) Enqueue(item T) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.items = append(q.items, item)
}

func (q *GenericQueue[T]) Dequeue() (T, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	var zero T
	if len(q.items) == 0 {
		return zero, errors.New("queue is empty")
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}

func RunQueue() {
	// Demonstrate BasicQueue
	fmt.Println("Basic Queue:")
	basicQueue := &BasicQueue{}
	basicQueue.Enqueue(1)
	basicQueue.Enqueue(2)
	basicQueue.Enqueue(3)

	front, _ := basicQueue.Front()
	fmt.Println("Front item:", front)

	dequeued, _ := basicQueue.Dequeue()
	fmt.Println("Dequeued item:", dequeued)
	fmt.Println("Queue size:", basicQueue.Size())

	// Demonstrate ConcurrentQueue
	fmt.Println("\nConcurrent Queue:")
	concurrentQueue := &ConcurrentQueue{}
	concurrentQueue.Enqueue(10)
	concurrentQueue.Enqueue(20)

	// Demonstrate Generic Queue
	fmt.Println("\nGeneric Queue:")
	stringQueue := &GenericQueue[string]{}
	stringQueue.Enqueue("Hello")
	stringQueue.Enqueue("World")
	fmt.Println(stringQueue.Dequeue())
	fmt.Println(stringQueue.Dequeue())

	intQueue := &GenericQueue[int]{}
	intQueue.Enqueue(100)
	intQueue.Enqueue(200)
	fmt.Println(intQueue.Dequeue())
	fmt.Println(intQueue.Dequeue())
}
