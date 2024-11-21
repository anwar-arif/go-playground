package data_structure

import "fmt"

// Node represents an element in the doubly linked list
type Node struct {
	Value interface{}
	Prev  *Node
	Next  *Node
}

// DoublyLinkedList is the main list structure
type DoublyLinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		Head: nil,
		Tail: nil,
		Size: 0,
	}
}

// InsertFront adds a new node at the beginning of the list
func (dll *DoublyLinkedList) InsertFront(value interface{}) {
	newNode := &Node{Value: value}

	if dll.Head == nil {
		dll.Head = newNode
		dll.Tail = newNode
	} else {
		newNode.Next = dll.Head
		dll.Head.Prev = newNode
		dll.Head = newNode
	}
	dll.Size++
}

// InsertBack adds a new node at the end of the list
func (dll *DoublyLinkedList) InsertBack(value interface{}) {
	newNode := &Node{Value: value}

	if dll.Tail == nil {
		dll.Head = newNode
		dll.Tail = newNode
	} else {
		newNode.Prev = dll.Tail
		dll.Tail.Next = newNode
		dll.Tail = newNode
	}
	dll.Size++
}

// RemoveFront removes the first node from the list
func (dll *DoublyLinkedList) RemoveFront() (interface{}, bool) {
	if dll.Head == nil {
		return nil, false
	}

	value := dll.Head.Value
	dll.Head = dll.Head.Next

	if dll.Head == nil {
		dll.Tail = nil
	} else {
		dll.Head.Prev = nil
	}

	dll.Size--
	return value, true
}

// RemoveBack removes the last node from the list
func (dll *DoublyLinkedList) RemoveBack() (interface{}, bool) {
	if dll.Tail == nil {
		return nil, false
	}

	value := dll.Tail.Value
	dll.Tail = dll.Tail.Prev

	if dll.Tail == nil {
		dll.Head = nil
	} else {
		dll.Tail.Next = nil
	}

	dll.Size--
	return value, true
}

func (dll *DoublyLinkedList) IsEmpty() bool {
	return dll.Size == 0
}

func RunDoublyLinkedList() {
	list := NewDoublyLinkedList()
	list.InsertBack(1)
	list.InsertBack(2)
	list.InsertBack(3)
	list.InsertBack(4)

	for !list.IsEmpty() {
		fmt.Println(list.RemoveBack())
	}

}
