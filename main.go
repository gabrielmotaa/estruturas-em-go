package main

import (
	"fmt"

	"github.com/gabrielmotaa/datastructures/stack"
)

func main() {
	names := []string{"Daniel", "Gabriel", "Pedro"}
	var s stack.Stack

	fmt.Println("Stack:")
	fmt.Printf("Before filling the stack: %v\n\n", s.String())

	for _, name := range names {
		fmt.Println("Adding:", name)
		s.Push(name)
	}

	fmt.Printf("\nAfter filling the stack: %v\n", s.String())

	fmt.Println("Emptying stack:")
	for !s.IsEmpty() {
		value, _ := s.Pop()
		fmt.Println("\t" + value)
	}
}
