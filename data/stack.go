package data

import "fmt"

type Stack struct {
	items []int
	size  int
}

func NewStack() *Stack {
	return &Stack{
		items: make([]int, 0),
		size:  0,
	}
}

func (s *Stack) Push(value int) {
	s.items = append(s.items, value)
	s.size++
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}
	value := s.items[s.size-1]
	s.items = s.items[:s.size-1]
	s.size--
	return value, nil
}

func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}
	return s.items[s.size-1], nil
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Len() int {
	return s.size
}

func (s *Stack) Items() []int {
	return append([]int{}, s.items...)
}
