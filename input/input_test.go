package input

import (
	"reflect"
	"testing"

	"push-swap/stack"
)

func TestParseArgs(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    stack.Stack
		wantErr bool
	}{
		{"une chaîne", []string{"2 1 3"}, stack.Stack{2, 1, 3}, false},
		{"plusieurs arguments", []string{"2", "1", "3"}, stack.Stack{2, 1, 3}, false},
		{"négatif", []string{"-4 0 7"}, stack.Stack{-4, 0, 7}, false},
		{"non entier", []string{"0 one 2"}, nil, true},
		{"doublon", []string{"1 2 1"}, nil, true},
		// À toi d'ajouter : aucun argument, argument vide,
		// "1.5", un nombre énorme (99999999999999999999), "+3"...
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseArgs(tt.args)
			if (err != nil) != tt.wantErr {
				t.Fatalf("erreur = %v, wantErr = %v", err, tt.wantErr)
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
