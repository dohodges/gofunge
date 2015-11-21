package gofunge

import (
	"fmt"
)

type Stack struct {
	data []rune
}

func NewStack() *Stack {
	return &Stack{
		data: make([]rune, 0),
	}
}

func (s *Stack) Depth() int {
	return len(s.data)
}

func (s *Stack) Push(value rune) {
	s.data = append(s.data, value)
}

func (s *Stack) Pop() (rune, error) {
	depth := s.Depth()
	if depth == 0 {
		return 0, fmt.Errorf("befungo.Stack: empty")
	}

	value := s.data[depth-1]
	s.data = s.data[:depth-1]

	return value, nil
}

func (s *Stack) Clear() {
	s.data = make([]rune, 0)
}
