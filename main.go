package main

import (
	"fmt"

	"github.com/gabrielmotaa/datastructures/queue"
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

	fmt.Println("\nQueue:")
	q := queue.New()
	fmt.Printf("Before filling the queue: %v\n\n", q.String())

	for _, name := range names {
		fmt.Println("Adding:", name)
		q.Push(name)
	}

	fmt.Printf("\nAfter filling the queue: %v\n", q.String())

	fmt.Println("Emptying queue:")
	for !q.IsEmpty() {
		value, _ := q.Pop()
		fmt.Println("\t" + value)
	}
}
