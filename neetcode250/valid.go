// an edge case i want to check is if the length of the string is less than or equal to 1, it is not a valid string
// return false; i want to declare a left pointer and a right pointer, such that the left pointer
// starts from the first element in the string and the right pointer starts from the last element
// of the string, based on the string characters e.g if an open bracket is at the left pointer and
// its corresponding closed bracket is not at the right pointer, return false; otherwise increment
// the length of the left pointer and decrement the length of the right pointer and check the string
// characters again based on open and closed brackets, this mechanism is done in a for loop such that
// the left pointer will always be less than right pointer, now if all open brackets has corresponding
// closed brackets, return true
package main

func isValid(s string) bool {
	if len(s) <= 1 {
		return false
	}

	L := 0
	R := len(s) - 1

	for L < R {
		if s[L] == '(' && s[R] == ')' {
			L++
			R--
		} else if s[L] == '{' && s[R] == '}' {
			L++
			R--
		} else if s[L] == '[' && s[R] == ']' {
			L++
			R--
		} else if s[L] == '(' && s[L+1] == ')' {
			L += 2
			R = L + 3
		} else if s[L] == '{' && s[L+1] == '}' {
			L += 2
			R = L + 3
		} else if s[L] == '[' && s[L+1] == ']' {
			L += 2
			R = L + 3
		} else {
			return false
		}
	}
	return true
}

// Time Complexity: O(n)
// Space Complexity: O(n)
type Stack struct {
	stack []rune
}

func Newstack(items []rune) *Stack {
	return &Stack{
		stack: items,
	}
}

func (s *Stack) Peek() rune {
	if len(s.stack) == 0 {
		return ' '
	}
	item := s.stack[len(s.stack)-1]
	return item
}

func (s *Stack) Push(item rune) {
	s.stack = append(s.stack, item)
}

func (s *Stack) Pop() rune {
	if len(s.stack) == 0 {
		return ' '
	}
	item := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return item
}

func (s *Stack) Empty() bool {
	if len(s.stack) == 0 {
		return true
	}
	return false
}

func isValid1(s string) bool {
	stack := Newstack([]rune{})
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stack.Push(v)
		} else if v == ')' {
			if !stack.Empty() && stack.Peek() == '(' {
				stack.Pop()
			} else {
				return false
			}
		} else if v == ']' {
			if !stack.Empty() && stack.Peek() == '[' {
				stack.Pop()
			} else {
				return false
			}
		} else if v == '}' {
			if !stack.Empty() && stack.Peek() == '{' {
				stack.Pop()
			} else {
				return false
			}
		}
	}
	return stack.Empty()
}
