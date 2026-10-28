package sorter

import (
	"testing"

	"push-swap/stack"
)

// check vérifie que les instructions renvoyées trient bien in.
func check(t *testing.T, in []int, maxOps int) {
	t.Helper()
	ops := Sort(append(stack.Stack{}, in...))

	m := stack.Machine{A: append(stack.Stack{}, in...)}
	for _, op := range ops {
		if err := m.Apply(op); err != nil {
			t.Fatalf("%v : %v", in, err)
		}
	}
	if !m.IsSorted() {
		t.Errorf("%v : %v ne trie pas (résultat %v)", in, ops, m.A)
	}
	if len(ops) > maxOps {
		t.Errorf("%v : %d instructions, maximum %d", in, len(ops), maxOps)
	}
}

func TestSort2(t *testing.T) {
	check(t, []int{1, 2}, 0)
	check(t, []int{2, 1}, 1)
}

func TestSort3(t *testing.T) {
	cases := [][]int{
		{1, 2, 3}, {2, 1, 3}, {1, 3, 2},
		{2, 3, 1}, {3, 1, 2}, {3, 2, 1},
	}
	for _, c := range cases {
		check(t, c, 2)
	}
}
