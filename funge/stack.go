package funge

type Stack struct {
	funge  Funge
	stacks [][]rune
}

func NewStack(funge Funge) *Stack {
	return &Stack{
		funge:  funge,
		stacks: make([][]rune, 1),
	}
}

func (s *Stack) Depth() int {
	return len(s.topOfStackStack())
}

func (s *Stack) Push(values ...rune) {
	for _, value := range values {
		s.setTopOfStackStack(append(s.topOfStackStack(), value))
	}
}

func (s *Stack) Pop() rune {
	depth := s.Depth()
	if depth == 0 {
		return 0
	}

	toss := s.topOfStackStack()
	value := toss[depth-1]
	s.setTopOfStackStack(toss[:depth-1])

	return value
}

func (s *Stack) PopVector() Vector {
	vector := NewVector(int(s.funge))
	for axis := Axis(s.funge - 1); axis >= XAxis; axis-- {
		value := int32(s.Pop())
		vector.Set(axis, value)
	}

	return vector
}

func (s *Stack) Clear() {
	s.setTopOfStackStack(make([]rune, 0))
}

func (s *Stack) topOfStack() int {
	return len(s.stacks) - 1
}

func (s *Stack) topOfStackStack() []rune {
	return s.stacks[s.topOfStack()]
}

func (s *Stack) setTopOfStackStack(stack []rune) {
	s.stacks[s.topOfStack()] = stack
}
