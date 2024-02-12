package stack

import "errors"

type Stack struct {
	data []int
}

func (s *Stack) Push(n int) {
	s.data = append(s.data, n)
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("no items in the stack")
	}

	i := s.Size() - 1
	p := s.data[i]
	s.data = s.data[:i]

	return p, nil
}

func (s Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("no items in the stack")
	}

	i := s.Size() - 1

	return s.data[i], nil
}

func (s Stack) Size() int {
	return len(s.data)
}

func (s Stack) IsEmpty() bool {
	return len(s.data) == 0
}
