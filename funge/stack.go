package funge

type Stack struct {
	stacks [][]rune
}

func NewStack() *Stack {
	return &Stack{
		stacks: make([][]rune, 1),
	}
}

func (s *Stack) Depth() int {
	return len(s.topOfStackStack())
}

func (s *Stack) Push(value rune) {
	s.setTopOfStackStack(append(s.topOfStackStack(), value))
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
