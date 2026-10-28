package sorter

import (
	"math/rand"
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
func permutations(n int) [][]int {
	var res [][]int
	used := make([]bool, n)
	var gen func(cur []int)
	gen = func(cur []int) {
		if len(cur) == n {
			res = append(res, append([]int{}, cur...))
			return
		}
		for v := 1; v <= n; v++ {
			if used[v-1] {
				continue
			}
			used[v-1] = true
			gen(append(cur, v))
			used[v-1] = false
		}
	}
	gen(nil)
	return res
}

func TestSort4And5(t *testing.T) {
	for _, n := range []int{4, 5} {
		for _, p := range permutations(n) {
			check(t, p, 12)
		}
	}
}

func TestSortLarge(t *testing.T) {
	r := rand.New(rand.NewSource(1))
	for _, n := range []int{6, 7, 10, 100, 500} {
		in := r.Perm(n)
		for i := range in {
			in[i] -= n / 2 // introduit des nombres négatifs
		}
		check(t, in, 100000)
		t.Logf("n=%d : %d instructions", n, len(Sort(append(stack.Stack{}, in...))))
	}
}
func TestAverageOps(t *testing.T) {
	r := rand.New(rand.NewSource(1))
	for _, n := range []int{6, 7, 8, 9, 10, 12, 15, 20, 30, 50, 100} {
		total := 0
		const runs = 50
		for i := 0; i < runs; i++ {
			in := r.Perm(n)
			check(t, in, 100000)
			total += len(Sort(append(stack.Stack{}, in...)))
		}
		t.Logf("n=%3d : moyenne %d instructions", n, total/runs)
	}
}
func TestOptimize(t *testing.T) {
	cases := []struct{ in, want []string }{
		{[]string{"ra", "rb"}, []string{"rr"}},
		{[]string{"ra", "rra"}, []string{}},
		{[]string{"pb", "pa"}, []string{}},
		{[]string{"ra", "ra", "rra", "rb"}, []string{"rr"}},
	}
	for _, c := range cases {
		got := Optimize(c.in)
		if len(got) != len(c.want) {
			t.Fatalf("%v : got %v, want %v", c.in, got, c.want)
		}
		for i := range got {
			if got[i] != c.want[i] {
				t.Errorf("%v : got %v, want %v", c.in, got, c.want)
			}
		}
	}
}

func solveSmall(a stack.Stack) []string {
	s := New(a)
	s.sortSmall()
	return s.Ops
}

func solveChunk(a stack.Stack, size int) []string {
	s := New(a)
	s.chunk(size)
	return s.Ops
}

// avg : nombre moyen d'instructions (après Optimize) sur runs piles aléatoires.
func avg(n, runs int, solve func(stack.Stack) []string) int {
	r := rand.New(rand.NewSource(1))
	total := 0
	for i := 0; i < runs; i++ {
		total += len(Optimize(solve(stack.Stack(r.Perm(n)))))
	}
	return total / runs
}

func TestTuneSmallLimit(t *testing.T) {
	for _, n := range []int{6, 8, 10, 12, 15, 20, 25, 30, 40} {
		small := avg(n, 50, solveSmall)
		chunks := avg(n, 50, func(a stack.Stack) []string { return solveChunk(a, chunkSize(n)) })
		t.Logf("n=%3d : petites=%4d  chunks=%4d", n, small, chunks)
	}
}

func TestTuneChunkSize(t *testing.T) {
	for _, size := range []int{10, 12, 14, 16, 18, 20, 22} {
		got := avg(100, 30, func(a stack.Stack) []string { return solveChunk(a, size) })
		t.Logf("n=100 size=%3d : %d instructions", size, got)
	}
	for _, size := range []int{10, 15, 20, 25, 30, 35, 40} {
		got := avg(500, 10, func(a stack.Stack) []string { return solveChunk(a, size) })
		t.Logf("n=500 size=%3d : %d instructions", size, got)
	}
}
