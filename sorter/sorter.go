package sorter

import (
	"sort"

	"push-swap/stack"
)

// smallLimit : jusqu'à cette taille, on utilise sortSmall ; au-delà, les chunks.
// À régler avec TestTuneSmallLimit.
const smallLimit = 30

// Sorter exécute des instructions sur une Machine et les mémorise.
type Sorter struct {
	M   stack.Machine
	Ops []string
}

// New copie a : l'appelant garde sa pile intacte.
func New(a stack.Stack) *Sorter {
	return &Sorter{M: stack.Machine{A: append(stack.Stack{}, a...)}}
}

// do applique une instruction et la mémorise.
func (s *Sorter) do(op string) {
	_ = s.M.Apply(op) // instructions écrites en dur : pas d'erreur possible
	s.Ops = append(s.Ops, op)
}

// Sort renvoie la liste d'instructions (nettoyée) qui trie a.
func Sort(a stack.Stack) []string {
	s := New(a)
	s.sort()
	return Optimize(s.Ops)
}

func (s *Sorter) sort() {
	if s.M.IsSorted() {
		return
	}
	switch n := len(s.M.A); {
	case n == 2:
		s.do("sa")
	case n == 3:
		s.sort3()
	case n <= smallLimit:
		s.sortSmall()
	default:
		s.chunk(chunkSize(n))
	}
}

// sort3 trie exactement 3 éléments en 2 instructions maximum.
func (s *Sorter) sort3() {
	a, b, c := s.M.A[0], s.M.A[1], s.M.A[2]
	switch {
	case a > b && a > c:
		s.do("ra")
	case b > a && b > c:
		s.do("rra")
	}
	// On relit la pile : a, b, c sont les valeurs d'avant les rotations.
	if s.M.A[0] > s.M.A[1] {
		s.do("sa")
	}
}

// minIndex renvoie la position du plus petit élément de a.
func (s *Sorter) minIndex() int {
	idx := 0
	for i, v := range s.M.A {
		if v < s.M.A[idx] {
			idx = i
		}
	}
	return idx
}

// bringMinToTop remonte le plus petit élément de a au sommet (ra ou rra).
func (s *Sorter) bringMinToTop() {
	i := s.minIndex()
	n := len(s.M.A)
	if i <= n/2 {
		for ; i > 0; i-- {
			s.do("ra")
		}
	} else {
		for ; i < n; i++ {
			s.do("rra")
		}
	}
}

// sortSmall : on envoie les plus petits dans b, on trie les 3 restants,
// puis on ramène tout avec pa. Valable pour n >= 3.
func (s *Sorter) sortSmall() {
	for len(s.M.A) > 3 {
		s.bringMinToTop()
		s.do("pb")
	}
	s.sort3()
	for len(s.M.B) > 0 {
		s.do("pa")
	}
}

// bringMaxOfBToTop remonte le plus grand élément de b au sommet (rb ou rrb).
func (s *Sorter) bringMaxOfBToTop() {
	idx := 0
	for i, v := range s.M.B {
		if v > s.M.B[idx] {
			idx = i
		}
	}
	n := len(s.M.B)
	if idx <= n/2 {
		for ; idx > 0; idx-- {
			s.do("rb")
		}
	} else {
		for ; idx < n; idx++ {
			s.do("rrb")
		}
	}
}

// toRanks remplace chaque valeur de a par son rang (0 = le plus petit).
// L'ordre relatif ne change pas : les instructions restent valables
// pour les vraies valeurs.
func (s *Sorter) toRanks() {
	sorted := append([]int{}, s.M.A...)
	sort.Ints(sorted)
	rank := make(map[int]int, len(sorted))
	for i, v := range sorted {
		rank[v] = i
	}
	for i, v := range s.M.A {
		s.M.A[i] = rank[v]
	}
}

// chunkSize choisit la taille des tranches (à régler avec TestTuneChunkSize).
func chunkSize(n int) int {
	size := n / 5
	if n > 100 {
		size = n / 11
	}
	if size < 1 {
		size = 1
	}
	return size
}

// chunk trie par tranches de rangs de taille size.
func (s *Sorter) chunk(size int) {
	s.toRanks()

	// Aller : tout envoyer dans b, tranche après tranche.
	pushed := 0
	for len(s.M.A) > 0 {
		r := s.M.A[0]
		switch {
		case r < pushed: // petit : on l'envoie tout en bas de b
			s.do("pb")
			if len(s.M.B) > 1 {
				s.do("rb")
			}
			pushed++
		case r < pushed+size: // dans la tranche en cours
			s.do("pb")
			pushed++
		default: // trop grand pour l'instant
			s.do("ra")
		}
	}

	// Retour : toujours ramener le plus grand de b.
	for len(s.M.B) > 0 {
		s.bringMaxOfBToTop()
		s.do("pa")
	}
}
