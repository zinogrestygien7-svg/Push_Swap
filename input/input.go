package input

import (
	"errors"
	"strconv"
	"strings"

	"push-swap/stack"
)

// ErrInvalid est renvoyée pour un argument non entier ou un doublon.
var ErrInvalid = errors.New("Error")

// ParseArgs transforme les arguments de la ligne de commande en pile.
// Le premier nombre se retrouve au sommet (index 0).
func ParseArgs(args []string) (stack.Stack, error) {
	fields := strings.Fields(strings.Join(args, " "))

	s := make(stack.Stack, 0, len(fields))
	seen := make(map[int]bool)

	for _, f := range fields {
		n, err := strconv.Atoi(f)
		if err != nil {
			return nil, ErrInvalid
		}
		if seen[n] {
			return nil, ErrInvalid
		}
		seen[n] = true
		s = append(s, n)
	}
	return s, nil
}
