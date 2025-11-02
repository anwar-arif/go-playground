package payment_processor

import (
	"container/heap"
	"fmt"
	"time"
)

type Payment struct {
	ID        string
	Priority  int
	Timestamp time.Time
}

type PaymentQueue struct {
	items []*Payment
}

func NewPaymentQueue() *PaymentQueue {
	pq := &PaymentQueue{
		items: make([]*Payment, 0),
	}
	heap.Init(pq)
	return pq
}

func (pq *PaymentQueue) Len() int {
	return len(pq.items)
}

func (pq *PaymentQueue) Less(i, j int) bool {
	if pq.items[i].Priority != pq.items[j].Priority {
		return pq.items[i].Priority > pq.items[j].Priority
	} else {
		return pq.items[i].Timestamp.Before(pq.items[j].Timestamp)
	}
}

func (pq *PaymentQueue) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
}

func (pq *PaymentQueue) Push(x any) {
	pq.items = append(pq.items, x.(*Payment))
}

func (pq *PaymentQueue) Pop() any {
	old := pq.items
	n := len(old)
	item := old[n-1]
	pq.items = old[:n-1]
	return item
}

func (pq *PaymentQueue) PushVal(payment *Payment) {
	heap.Push(pq, payment)
}

func (pq *PaymentQueue) PopVal() *Payment {
	payment := heap.Pop(pq).(*Payment)
	return payment
}

func (pq *PaymentQueue) Top() *Payment {
	if len(pq.items) == 0 {
		return nil
	}
	return pq.items[0]
}

func RunPaymentProcessor() {
	pq := NewPaymentQueue()

	pq.PushVal(&Payment{ID: "A", Priority: 2, Timestamp: time.Now().Add(10 * time.Second)})
	pq.PushVal(&Payment{ID: "B", Priority: 3, Timestamp: time.Now().Add(1 * time.Minute)})
	pq.PushVal(&Payment{ID: "C", Priority: 3, Timestamp: time.Now().Add(5 * time.Second)})

	p := pq.Top()
	fmt.Printf("Top: %s (Priority: %d, Time: %v)\n", p.ID, p.Priority, p.Timestamp)
	for pq.Len() > 0 {
		p := pq.PopVal()
		fmt.Printf("Processed: %s (Priority: %d, Time: %v)\n", p.ID, p.Priority, p.Timestamp)
	}

	fmt.Printf("time: %v\n", time.Now().Format("Mon Jan 2 03:04:05 PM MST"))
}
