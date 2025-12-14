package main

import (
	"fmt"
)

type Stack struct {
	stack []int // a stack is an array with three restrictions
}

// only the last element of a stack can be read
func (s *Stack) Read() int {
	if len(s.stack) == 0 {
		fmt.Println("Error: Stack is empty")
		return 0
	}
	res := s.stack[len(s.stack)-1]
	return res
}

// data can be inserted only at the end of a stack
func (s *Stack) Push(n int) {
	s.stack = append(s.stack, n)
}

// removing an element from the end/top of the stack
func (s *Stack) Pop() int {
	if len(s.stack) == 0 {
		fmt.Println("Error: Stack is empty")
		return 0
	}
	res := s.stack[len(s.stack)-1]
	// data can be deleted only from the end of a stack
	s.stack = s.stack[:len(s.stack)-1]

	return res
}

// size of the stack
func (s *Stack) Size() int {
	return len(s.stack)
}
