package main

import "fmt"

func inorder(root *TreeNode) {
	stack := []*TreeNode{}
	curr := root

	for curr != nil || len(stack) > 0 {
		if curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		} else {
			curr = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			fmt.Println(curr.Val)
			curr = curr.Right
		}
	}
}

func inorder1(root *TreeNode) {
	stack := []*TreeNode{}
	curr := root

	for curr != nil || len(stack) > 0 {
		if curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		} else {
			curr = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			fmt.Println(curr.Val)
			curr = curr.Right
		}
	}
}

func inorder2(root *TreeNode) {
	stack := []*TreeNode{}
	curr := root

	for curr != nil || len(stack) > 0 {
		if curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		} else {
			curr = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			fmt.Println(curr.Val)
			curr = curr.Right
		}
	}
}

func inorder3(root *TreeNode) {
	stack := []*TreeNode{}
	curr := root

	for curr != nil || len(stack) > 0 {
		if curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		} else {
			curr = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			fmt.Println(curr.Val)
			curr = curr.Right
		}
	}

}

func inorder4(root *TreeNode) {
	stack := []*TreeNode{}
	curr := root

	for curr != nil && len(stack) > 0 {
		if curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		} else {
			curr = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			fmt.Println(curr.Val)
			curr = curr.Right
		}

	}
}

func inorder5(root *TreeNode) {
	stack := []*TreeNode{}
	curr := root

	for curr != nil || len(stack) > 0 {
		if curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		} else {
			curr = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			fmt.Println(curr.Val)
			curr = curr.Right
		}
	}
}

func inorder6(root *TreeNode) {
	stack := []*TreeNode{}
	curr := root

	for curr != nil || len(stack) > 0 {
		if curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		} else {
			curr = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			fmt.Println(curr.Val)
			curr = curr.Right
		}
	}
}

func inorder7(root *TreeNode) {
	stack := []*TreeNode
	curr := root

	for curr != nil || len(stack) > 0 {
		if curr != nil {
			curr = append(stack, curr)
			curr = curr.Left
		} else {
			curr = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			fmt.Println(curr.Val)
			curr = curr.Right
		}
	}
}

func inorder8(root *TreeNode) {
	stack := []*TreeNode{}
	curr := root

	for curr != nil || len(stack) > 0 {
		if curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		} else {
			curr = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			fmt.Println(curr.Val)
			curr = curr.Right
		}
	}
}

func inorder9(root *TreeNode) {
	stack := []*TreeNode{}
	curr := root

	for curr != nil || len(stack) > 0 {
		if curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		} else {
			curr = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			fmt.Println(curr.Val)
			curr = curr.Right
		}
	}
}

func inorder10(root *TreeNode) {
	stack := []*TreeNode{}
	curr := root

	for curr != nil || len(stack) > 0 {
		if curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		} else {
			curr = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			fmt.Println(curr.Val)
			curr = curr.Right
		}
	}
}

func inorder11(root *TreeNode) {
	stack := []*TreeNode{}
	curr := root

	for curr != nil || len(stack) > 0 {
		if curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		} else {
			curr = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			fmt.Println(curr.Val)
			curr = curr.Right
		}
	}

}

func inorder12(root *TreeNode) {
	stack := []*TreeNode{}
	curr := root

	for curr != nil || len(stack) > 0 {
		if curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		} else {
			curr = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			fmt.Println(curr.Val)
			curr = curr.Right
		}
	}
}

func inorder13(root *TreeNode) {
	stack := []*TreeNode{}
	curr := root

	for curr != nil || len(stack) > 0 {
		if curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		} else {
			curr = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			fmt.Println(curr.Val)
			curr = curr.Right
		}
	}
}
