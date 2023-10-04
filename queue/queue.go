package queue

import (
	"errors"

	"github.com/gabrielmotaa/datastructures/stack"
)

// A implementação de uma fila usando stacks
// é mais simples mas toda operação envolve
// na fila esvaziar uma stack e preencher outra.

type Queue struct {
	inputStack  *stack.Stack
	outputStack *stack.Stack
}

func (q *Queue) Push(value string) {
	for !q.outputStack.IsEmpty() {
		outputValue, _ := q.outputStack.Pop()
		q.inputStack.Push(outputValue)
	}
	q.inputStack.Push(value)
}

func (q *Queue) Pop() (string, error) {
	if q.IsEmpty() {
		return "", errors.New("empty queue")
	}

	for !q.inputStack.IsEmpty() {
		inputValue, _ := q.inputStack.Pop()
		q.outputStack.Push(inputValue)
	}

	return q.outputStack.Pop()
}

func (q *Queue) IsEmpty() bool {
	return q.inputStack.IsEmpty() && q.outputStack.IsEmpty()
}

func (q *Queue) String() string {
	// Queremos mostrar a lista na ordem correta,
	// ou seja, como ficaria na stack de output.
	if q.outputStack.IsEmpty() {
		for !q.inputStack.IsEmpty() {
			value, _ := q.inputStack.Pop()
			q.outputStack.Push(value)
		}
	}

	return q.outputStack.String()
}

func New() Queue {
	return Queue{
		inputStack:  &stack.Stack{},
		outputStack: &stack.Stack{},
	}
}
