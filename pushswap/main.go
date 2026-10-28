package main

import (
	"fmt"
	"os"

	"push-swap/input"
	"push-swap/sorter"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}

	a, err := input.ParseArgs(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	for _, op := range sorter.Sort(a) {
		fmt.Println(op)
	}
}
