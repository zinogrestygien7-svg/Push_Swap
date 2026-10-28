package sorter

import (
	"push-swap/stack"
)

type Sorter struct {
	M   stack.Machine
	Ops []string
}

func New(a stack.Stack) *Sorter {
	return &Sorter{M: stack.Machine{A: a}}
}
func (s *Sorter) do(op string) {
	_ = s.M.Apply(op)
	s.Ops = append(s.Ops, op)
}
func Sort(a stack.Stack) []string {
	s := New(a)
	s.sort()
	return s.Ops
}
func (s *Sorter) sort() {
	if s.M.IsSorted() {
		return
	}
	switch len(s.M.A) {
	case 2:
		s.do("sa")
	case 3:
		s.sort3()
	default:
		// a venir
	}
}
func (s *Sorter) sort3() {
	a, b, c := s.M.A[0], s.M.A[1], s.M.A[2]
	switch {
	case a > b && a > c:
		s.do("ra")
	case b > a && b > c:
		s.do("rra")
	}
	if s.M.A[0] > s.M.A[1] {
		s.do("sa")
	}
}
