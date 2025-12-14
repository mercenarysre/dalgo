package main

import (
	"fmt"
)

// Algorithm to traverse trees and graphs
// Pick a direction, say left, and keep following pointers as far down left as we can go until we reach null.
// Once we reach null, we backtrack to the parent node and then go right.
// We keep doing this until we have visited every node in the tree.
// This is the essence of depth-first search.

// Time Complexity: O(n), n is the number of nodes in the tree,
// regardless of the height of the tree.
// Space Complexity: O(h), where h is the height of the tree, which would be
//	O(logn)	for a balanced binary tree or O(n) for a skewed tree.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// TreeNode constructor
func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

// InOrder traversal
func (root *TreeNode) InOrder() {
	// base case
	if root == nil {
		return
	}
	root.Left.InOrder()
	fmt.Println(root.Val)
	root.Right.InOrder()
}

// PreOrder traversal
func (root *TreeNode) PreOrder() {
	// base case
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	root.Left.PreOrder()
	root.Right.PreOrder()
}

// PostOrder traversal
func (root *TreeNode) PostOrder() {
	// base case
	if root == nil {
		return
	}
	root.Left.PostOrder()
	root.Right.PostOrder()
	fmt.Println(root.Val)
}
