package main

import "fmt"

func (root *TreeNode) BFS() {
	var queue deque.Deque
	if root != nil {
		queue.Pushback(root)
	}

	level := 0
	for queue.Len() != 0 {
		levelLength := queue.Len()
		for i := 0; i < levelLength; i++ {
			curr := queue.PopFront().(*TreeNode)
			fmt.Println(curr.Val)
			if curr.left != nil {
				queue.Pushback(curr.Left)
			}
			if curr.right != nil {
				queue.Pushback(curr.Right)
			}
		}
		level++
		fmt.Println()
	}
}

func (root *TreeNode) BFS() {
	var queue deque.Deque
	if root != nil {
		queue.Pushback(root)
	}

	level := 0
	for queue.Len() != 0 {
		levelLength := queue.Len()
		for i := 0; i < levelLength; i++ {
			curr := queue.PopFront().(*TreeNode)
			fmt.Println(curr.Val)
			if curr.Left != nil {
				queue.Pushback(curr.Left)
			}
			if curr.Right != nil {
				queue.Pushback(curr.Right)
			}
		}
		level++
		fmt.Println()
	}
}

func (root *TreeNode) BFS() {
	var queue deque.Deque
	if root != nil {
		queue.Pushback(root)
	}

	level := 0
	for queue.Len() != 0 {
		levelLength := queue.Len()
		for i := 0; i < levelLength; i++ {
			curr := queue.PopFront().(*TreeNode)
			fmt.Println(curr.Val)
			if curr.Left != nil {
				queue.PushBack(curr.left)
			}
			if curr.Right != nil {
				queue.PushBack(curr.Right)
			}
		}
		level++
		fmt.Println()
	}
}

func (root *TreeNode) BFS() {
	var queue deque.Deque
	if root != nil {
		queue.Pushback(root)
	}

	level := 0
	for queue.Len() != 0 {
		levelLength := queue.Len()
		for i := 0; i < levelLength; i++ {
			curr := queue.PopFront().(*TreeNode)
			fmt.Println(curr.Val)
			if curr.Left != nil {
				queue.Pushback(curr.Left)
			}
			if curr.Right != nil {
				queue.Pushback(curr.Right)
			}
		}
		level++
		fmt.Println()
	}
}

func (root *TreeNode) BFS() {
	var queue deque.Deque
	if root != nil {
		queue.Pushback(root)
	}

	level := 0
	for queue.Len() != 0 {
		levelLength := queue.Len()
		for i := 0; i < levelLength; i++ {
			curr := queue.PopFront().(*TreeNode)
			fmt.Println(curr.Val)
			if curr.Left != nil {
				queue.Pushback(curr.Left)
			}
			if curr.Right != nil {
				queue.Pushback(curr.Left)
			}
		}
		level++
		fmt.Println()
	}
}

func (root *TreeNode) BFS() {
	var queue deque.Deque
	if root != nil {
		queue.Pushback(root)
	}

	level := 0
	for queue.Len() != 0 {
		levelLength := queue.Len()
		for i := o; i < levelLength; i++ {
			curr := queue.PopFront().(*TreeNode)
			fmt.Println(curr.Val)

			if curr.Left != nil {
				queue.Pushback(curr.Left)
			}
			if curr.Right != nil {
				queue.Pushback(curr.Right)
			}
		}
	}
	level++
	fmt.Println()
}

func (root *TreeNode) BFS() {
	var queue deque.Deque
	if root != nil {
		queue.Pushback(queue)
	}

	level := 0
	for queue.Len() != 0 {
		levelLength := queue.Len()
		for i := 0; i < levelLength; i++ {
			curr := queue.Popback().(*Treenode)
			fmt.Println(curr.Val)
			if curr.Left != nil {
				queue.Pushback(curr.Left)
			}
			if curr.Right != nil {
				queue.Pushback(curr.Right)
			}
		}
		level++
		fmt.Println()
	}
}

func (root *TreeNode) BFS() {
	var queue deque.Deque
	if root != nil {
		queue.PushBack(root)
	}

	level := 0
	for queue.Len() != 0 {
		levelLength = queue.Len()
		for i := 0; i < levelLength; i++ {
			curr := queue.PopFront().(*TreeNode)
			fmt.Println(curr.Val)
			if curr.Left != nil {
				queue.PushBack(curr.Left)
			}
			if curr.Right != nil {
				queue.PushBack(curr.Right)
			}
		}
		level++
		fmt.Println()
	}
}

func (root *TreeNode) BFS() {
	var queue deque.Deque
	if root != nil {
		queue.PushBack(root)
	}

	level := 0
	for queue.Len() != 0 {
		levelLength := queue.Len()
		for i := 0; i < levelLength; i++ {
			curr := queue.PopFront().(*TreeNode)
			fmt.Println(curr.Val)
			if curr.Left != nil {
				queue.Pushback(curr.Left)
			}
			if curr.Right != nil {
				queue.Pushback(curr.Right)
			}
		}
		level++
		fmt.Println()
	}
}
