package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Iterative Inorder Traversal
// Starts at Root, Left to Root to Right
// Time Complexity: O(n), doing O(1) work at each node
// Space Complexity: O(n), in the worst case, we have all of the nodes in the stack
func inorder(root *TreeNode) {
	stack := []*TreeNode{}
	curr := root

	for curr != nil || len(stack) > 0 {
		if curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		} else {
			// pop
			curr = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			fmt.Println(curr.Val)
			curr = curr.Right
		}
	}
}

// Iterative Preorder Traversal
// Starts at Root, Root	to Left	to Right
// Time Complexity: O(n), doing O(1) work at each node
// Space Complexity: O(n), in the worst case, we have all of the nodes in the stack
func preorder(root *TreeNode) {
	stack := []*TreeNode{}
	curr := root

	for curr != nil || len(stack) > 0 {
		if curr != nil {
			fmt.Println(curr.Val)
			if curr.Right != nil {
				stack = append(stack, curr.Right)
			}
			curr = curr.Left
		} else {
			// pop
			curr = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
	}
}

// Iterative Postorder Traversal
// Starts at Root, Left	to Right to Root
// Time Complexity: O(n), doing O(1) work at each node
// Space Complexity: O(n), in the worst case, we have all of the nodes in the stack
func postorder(root *TreeNode) {
	stack := []*TreeNode{root}
	visited := []bool{false}

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		v := visited[len(visited)-1]
		stack = stack[:len(stack)-1]
		visited = visited[:len(visited)-1]

		if curr != nil {
			if v {
				fmt.Println(curr.Val)
			} else {
				// push current again with visited = true
				stack = append(stack, curr)
				visited = append(visited, true)

				// push right then left
				stack = append(stack, curr.Right)
				visited = append(visited, false)

				stack = append(stack, curr.Left)
				visited = append(visited, false)
			}
		}
	}
}
