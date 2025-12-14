package main

import (
	"fmt"

	"github.com/gammazero/deque"
)

// A binary tree is a tree in which each node has
// zero, one or two children
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

// Search
func (root *TreeNode) Search(target int) bool {
	if root == nil {
		return false
	}

	if target > root.Val {
		return root.Right.Search(target)
	} else if target < root.Val {
		return root.Left.Search(target)
	}
	return true
}

// Depth First Search
func (root *TreeNode) PreOrder() {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	root.Left.PreOrder()
	root.Right.PreOrder()
}

func (root *TreeNode) InOrder() {
	if root == nil {
		return
	}
	root.Left.InOrder()
	fmt.Println(root.Val)
	root.Right.InOrder()
}

func (root *TreeNode) PostOrder() {
	if root == nil {
		return
	}
	root.Left.PostOrder()
	root.Right.PostOrder()
	fmt.Println(root.Val)
}

// Breadth First Search
func (root *TreeNode) BFS() {
	var queue deque.Deque
	if root != nil {
		queue.PushBack(root)
	}
	level := 0
	for queue.Len() != 0 {
		fmt.Printf("level %d: ", level)
		levelLength := queue.Len()
		for i := 0; i < levelLength; i++ {
			curr := queue.PopFront().(*TreeNode)
			fmt.Printf("%d ", curr.val)
			if curr.left != nil {
				queue.PushBack(curr.left)
			}
			if curr.right != nil {
				queue.PushBack(curr.right)
			}
		}
		level++
		fmt.Println()
	}
}

// root := NewTreeNode(10)
// root.Left = NewTreeNode(5)
// root.Right = NewTreeNode(15)
// root.Left.Left = NewTreeNode(3)
