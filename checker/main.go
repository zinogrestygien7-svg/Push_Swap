package main

import (
	"bufio"
	"fmt"
	"os"

	"push-swap/input"
	"push-swap/stack"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}

	a, err := input.ParseArgs(args)
	if err != nil {
		fail()
	}
	m := stack.Machine{A: a}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if err := m.Apply(line); err != nil {
			fail()
		}
	}
	if scanner.Err() != nil {
		fail()
	}

	if m.IsSorted() {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}

// fail affiche Error sur la sortie d'erreur et arrête le programme.
func fail() {
	fmt.Fprintln(os.Stderr, "Error")
	os.Exit(1)
}
