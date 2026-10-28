package stack

import (
	"reflect"
	"testing"
)

func TestSujetExemple(t *testing.T) {
	m := Machine{A: Stack{2, 1, 3, 6, 8, 5}}
	instrs := []string{"sa", "pb", "pb", "pb", "rb", "rrr", "pa", "pa", "pa"}

	for _, ins := range instrs {
		if err := m.Apply(ins); err != nil {
			t.Fatalf("erreur sur %q : %v", ins, err)
		}
	}

	want := Stack{1, 2, 3, 5, 6, 8}
	if !reflect.DeepEqual(m.A, want) {
		t.Errorf("a : got %v, want %v", m.A, want)
	}
	if len(m.B) != 0 {
		t.Errorf("b devrait être vide, got %v", m.B)
	}
	if !m.IsSorted() {
		t.Error("IsSorted devrait renvoyer true")
	}
}

func TestApplyInconnue(t *testing.T) {
	m := Machine{A: Stack{1, 2}}
	if err := m.Apply("xx"); err == nil {
		t.Error("une instruction inconnue doit renvoyer une erreur")
	}
}
