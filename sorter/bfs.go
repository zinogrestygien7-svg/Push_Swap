package sorter

import "push-swap/stack"

var allOps = []string{"sa", "sb", "ss", "pa", "pb", "ra", "rb", "rr", "rra", "rrb", "rrr"}

// node : un état des piles, avec le chemin qui y mène.
type node struct {
	m    stack.Machine
	prev int    // index du nœud précédent dans la file (-1 pour le départ)
	op   string // instruction qui a mené à cet état
}

func clone(m stack.Machine) stack.Machine {
	return stack.Machine{
		A: append(stack.Stack{}, m.A...),
		B: append(stack.Stack{}, m.B...),
	}
}

// key transforme un état en texte unique, pour savoir si on l'a déjà vu.
func key(m stack.Machine) string {
	b := make([]byte, 0, len(m.A)+len(m.B)+1)
	for _, v := range m.A {
		b = append(b, byte(v))
	}
	b = append(b, 255) // séparateur entre a et b
	for _, v := range m.B {
		b = append(b, byte(v))
	}
	return string(b)
}

// bfs trouve la plus courte suite d'instructions qui trie a.
// À réserver aux petites piles : le nombre d'états explose avec n.
func (s *Sorter) bfs() {
	s.toRanks() // valeurs 0..n-1 : tient dans un octet pour la clé

	queue := []node{{m: clone(s.M), prev: -1}}
	seen := map[string]bool{key(s.M): true}

	for head := 0; head < len(queue); head++ {
		if queue[head].m.IsSorted() {
			// On remonte le chemin à l'envers, puis on le rejoue dans le bon sens.
			var path []string
			for i := head; queue[i].prev != -1; i = queue[i].prev {
				path = append(path, queue[i].op)
			}
			for i := len(path) - 1; i >= 0; i-- {
				s.do(path[i])
			}
			return
		}
		for _, op := range allOps {
			next := clone(queue[head].m)
			_ = next.Apply(op)
			k := key(next)
			if !seen[k] {
				seen[k] = true
				queue = append(queue, node{m: next, prev: head, op: op})
			}
		}
	}
}
