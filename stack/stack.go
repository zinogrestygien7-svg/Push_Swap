package stack

type Stack []int

func (s *Stack) Swap() {
	if len(*s) < 2 {
		return
	}
	(*s)[0], (*s)[1] = (*s)[1], (*s)[0]
}

func (s *Stack) Rotation() {
	if len(*s) < 2 {
		return
	}
	first := (*s)[0]
	copy(*s, (*s)[1:])
	(*s)[len(*s)-1] = first
}
func (s *Stack) InvRotation() {
	if len(*s) < 2 {
		return
	}
	last := (*s)[len(*s)-1]
	copy((*s)[1:], *s)
	(*s)[0] = last
}
func Push(from, to *Stack) {
	if len(*from) == 0 {
		return
	}
	v := (*from)[0]
	*from = (*from)[1:]
	*to = append(Stack{v}, *to...)
}
