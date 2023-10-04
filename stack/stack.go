package stack

import (
	"errors"
	"strings"
)

// para não ter que lidar com genéricos é mais fácil
// escolher um tipo, no caso string

type node struct {
	value string
	next  *node
}

type Stack struct {
	head  *node
	count int
}

func (s *Stack) Push(value string) {
	s.head = &node{value, s.head}
	s.count++
}

func (s *Stack) IsEmpty() bool {
	return s.count == 0
}

func (s *Stack) Pop() (string, error) {
	if s.IsEmpty() {
		return "", errors.New("stack is empty")
	}

	value := s.head.value
	s.head = s.head.next
	s.count--

	return value, nil
}

func (s *Stack) String() string {
	if s.IsEmpty() {
		return "[]"
	}

	var sb strings.Builder
	i := 0
	sb.WriteString("[")

	current := s.head

	for current != nil {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(current.value)

		current = current.next
		i++
	}
	sb.WriteString("]")

	return sb.String()
}
