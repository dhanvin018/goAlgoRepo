package dataStructures

import fmt "fmt"

type Stack struct {
	items []int
	top   int
}

func NewStack(capacity int) *Stack {
	if capacity < 1 {
		return nil
	}
	return &Stack{
		items: make([]int, 0, capacity),
		top:   0,
	}
}

func (s *Stack) Push(value int) {
	s.items = append(s.items, value)
	s.top += 1
}

func (s *Stack) Pop() (int, error) {
	if s.top == 0 {
		return 0, fmt.Errorf("stack is empty")
	}
	s.top -= 1
	return s.items[s.top], nil
}
