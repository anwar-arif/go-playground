package data_structure

import "fmt"

type Stack struct {
	items []interface{}
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() interface{} {
	if len(s.items) == 0 {
		return nil
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *Stack) Peek() interface{} {
	if len(s.items) == 0 {
		return nil
	}
	return s.items[len(s.items)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}

func RunStack() {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	for s.Size() > 0 {
		item := s.Pop()
		fmt.Println(item)
	}
}
