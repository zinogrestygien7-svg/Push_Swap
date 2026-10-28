package stack

import "fmt"

// Machine contient les deux piles du sujet.
type Machine struct {
	A, B Stack
}

// --- Push ---
func (m *Machine) Pa() { Push(&m.B, &m.A) } // sommet de b -> sommet de a
func (m *Machine) Pb() { Push(&m.A, &m.B) } // sommet de a -> sommet de b

// --- Swap ---
func (m *Machine) Sa() { m.A.Swap() }
func (m *Machine) Sb() { m.B.Swap() }
func (m *Machine) Ss() { m.Sa(); m.Sb() }

// --- Rotate ---
func (m *Machine) Ra() { m.A.Rotation() }
func (m *Machine) Rb() { m.B.Rotation() }
func (m *Machine) Rr() { m.Ra(); m.Rb() }

// --- Reverse rotate ---
func (m *Machine) Rra() { m.A.InvRotation() }
func (m *Machine) Rrb() { m.B.InvRotation() }
func (m *Machine) Rrr() { m.Rra(); m.Rrb() }

// Apply exécute une instruction donnée sous forme de texte ("pa", "rra"...).
// Elle renvoie une erreur si l'instruction n'existe pas.
func (m *Machine) Apply(instr string) error {
	switch instr {
	case "pa":
		m.Pa()
	case "pb":
		m.Pb()
	case "sa":
		m.Sa()
	case "sb":
		m.Sb()
	case "ss":
		m.Ss()
	case "ra":
		m.Ra()
	case "rb":
		m.Rb()
	case "rr":
		m.Rr()
	case "rra":
		m.Rra()
	case "rrb":
		m.Rrb()
	case "rrr":
		m.Rrr()
	default:
		return fmt.Errorf("instruction inconnue : %q", instr)
	}
	return nil
}

// IsSorted vaut true si b est vide et si a est trié par ordre croissant.
func (m *Machine) IsSorted() bool {
	if len(m.B) != 0 {
		return false
	}
	for i := 1; i < len(m.A); i++ {
		if m.A[i-1] > m.A[i] {
			return false
		}
	}
	return true
}
