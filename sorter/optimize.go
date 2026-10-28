package sorter

// Optimize supprime les instructions qui s'annulent et fusionne celles
// qui peuvent l'être, jusqu'à ce que la liste ne change plus.
func Optimize(ops []string) []string {
	for {
		next := pass(ops)
		if len(next) == len(ops) {
			return next
		}
		ops = next
	}
}

var cancelPairs = map[[2]string]bool{
	{"ra", "rra"}: true, {"rra", "ra"}: true,
	{"rb", "rrb"}: true, {"rrb", "rb"}: true,
	{"rr", "rrr"}: true, {"rrr", "rr"}: true,
	{"pa", "pb"}: true, {"pb", "pa"}: true,
	{"sa", "sa"}: true, {"sb", "sb"}: true, {"ss", "ss"}: true,
}

var mergePairs = map[[2]string]string{
	{"sa", "sb"}: "ss", {"sb", "sa"}: "ss",
	{"ra", "rb"}: "rr", {"rb", "ra"}: "rr",
	{"rra", "rrb"}: "rrr", {"rrb", "rra"}: "rrr",
}

// pass fait un seul passage : on compare chaque instruction à la dernière
// conservée.
func pass(ops []string) []string {
	out := make([]string, 0, len(ops))
	for _, op := range ops {
		if n := len(out); n > 0 {
			pair := [2]string{out[n-1], op}
			if cancelPairs[pair] {
				out = out[:n-1]
				continue
			}
			if merged, ok := mergePairs[pair]; ok {
				out[n-1] = merged
				continue
			}
		}
		out = append(out, op)
	}
	return out
}
