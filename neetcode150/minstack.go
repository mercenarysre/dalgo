package main

import "fmt"

type MinStack struct {
	stack []int
}

func Constructor(items []int) *MinStack {
	return &MinStack{
		stack: items,
	}
}

// Time Complexity: O(1)
// Space Complexity: O(1)
func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
}

// Time Complexity: O(1)
// Space Complexity: O(1)
func (this *MinStack) Pop() int {
	if len(this.stack) == 0 {
		fmt.Println("Stack is Empty")
		return -1
	}
	res := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	return res
}

// Time Complexity: O(1)
// Space Complexity: O(1)
func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		fmt.Println("Stack is Empty")
		return -1
	}

	res := this.stack[len(this.stack)-1]
	return res
}

// Time Complexity: O(1)
// Space Complexity: O(1)
func (this *MinStack) Empty() bool {
	if len(this.stack) == 0 {
		return true
	}
	return false
}

// Time Complexity: O(n)
// Space Complexity: O(n)
func (this *MinStack) GetMin() int {

	if this.Empty() {
		fmt.Println("Stack is Empty")
		return -1
	}

	var tmp []int
	tmp1 := Constructor(tmp)

	min := this.Top()

	for !this.Empty() {
		val := this.Pop()
		if min <= val {
			min = min
		} else {
			min = val
		}
		/*
			if val < min {
			min = val
			}
		*/
		tmp1.Push(val)
	}

	for !tmp1.Empty() {
		val := tmp1.Pop()
		this.Push(val)
	}

	return min
}
