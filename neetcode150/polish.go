package main

import "strconv"

type Stack struct {
	stack []int
}

func NewStack(items []int) *Stack {
	return &Stack{
		stack: items,
	}
}

func (s *Stack) Empty() bool {
	if len(s.stack) == 0 {
		return true
	}
	return false
}

func (s *Stack) Read() int {
	if s.Empty() {
		return -1
	}

	res := s.stack[len(s.stack)-1]
	return res
}

func (s *Stack) Pop() int {
	if s.Empty() {
		return -1
	}

	res := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return res
}

func (s *Stack) Push(n int) {
	s.stack = append(s.stack, n)
}

// Time Complexity: O(n)
// Space Complexity: O(n)
func evalRPN(tokens []string) int {
	var items []int
	stacks := NewStack(items)

	for _, token := range tokens {
		if token == "+" {
			a := stacks.Pop()
			b := stacks.Pop()
			stacks.Push(b + a)
			// stacks.stack = append(stacks.stack, b+a)
		} else if token == "-" {
			a := stacks.Pop()
			b := stacks.Pop()
			stacks.Push(b - a)
			// stacks.stack = append(stacks.stack, b-a)
		} else if token == "*" {
			a := stacks.Pop()
			b := stacks.Pop()
			stacks.Push(b * a)
			// stacks.stack = append(stacks.stack, b*a)
		} else if token == "/" {
			a := stacks.Pop()
			b := stacks.Pop()
			stacks.Push(b / a)
			// stacks.stack = append(stacks.stack, b/a)
		} else {
			num, _ := strconv.Atoi(token)
			stacks.Push(num)
			// stacks.stack = append(stacks.stack, num)
		}
	}
	return stacks.Read()
}
