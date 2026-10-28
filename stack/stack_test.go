package stack

import (
	"reflect"
	"testing"
)

func TestRotation(t *testing.T) {
	s := Stack{1, 2, 3}
	s.Rotation()
	want := Stack{2, 3, 1}
	if !reflect.DeepEqual(s, want) {
		t.Errorf("Got %v, want %v", s, want)
	}
}
func TestSwap(t *testing.T) {
	s := Stack{1, 2, 3}
	s.Swap()
	want := Stack{2, 1, 3}
	if !reflect.DeepEqual(s, want) {
		t.Errorf("Got %v, want %v", s, want)
	}
}
func TestInvRotation(t *testing.T) {
	s := Stack{1, 2, 3}
	s.InvRotation()
	want := Stack{3, 1, 2}
	if !reflect.DeepEqual(s, want) {
		t.Errorf("Got %v, want %v", s, want)
	}
}
func TestPush(t *testing.T) {
	from := Stack{1, 2, 3}
	to := Stack{4, 5}

	Push(&from, &to)

	wantFrom := Stack{2, 3}
	wantTo := Stack{1, 4, 5}

	if !reflect.DeepEqual(from, wantFrom) {
		t.Errorf("from: got %v, want %v", from, wantFrom)
	}
	if !reflect.DeepEqual(to, wantTo) {
		t.Errorf("to: got %v, want %v", to, wantTo)
	}
}
func TestPushEmpty(t *testing.T) {
	from := Stack{}
	to := Stack{4, 5}

	Push(&from, &to)

	if len(from) != 0 || !reflect.DeepEqual(to, Stack{4, 5}) {
		t.Errorf("push depuis une pile vide ne doit rien changer, got from=%v to=%v", from, to)
	}
}
func TestSwapSingle(t *testing.T) {
	s := Stack{1}
	s.Swap()
	if !reflect.DeepEqual(s, Stack{1}) {
		t.Errorf("got %v, want [1]", s)
	}
}
