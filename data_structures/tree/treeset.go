package main

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

type TreeSet struct {
	Root *TreeNode
}

// Time Complexity: O(logn)
// Space Complexity: O(h)
func (root *TreeNode) Search(val int) bool {
	if root == nil {
		return false
	}

	if val < root.Val {
		root.Left = root.Left.Search(val)
	} else if val > root.Val {
		root.Right = root.Right.Search(val)
	}
	return true
}

// Time Complexity: O(logn)
// Space Complexity: O(h)
func (root *TreeNode) Insert(val int) *TreeNode {
	if root == nil {
		return NewTreeNode(val)
	}

	if val < root.Val {
		root.Left = root.Left.Insert(val)
	} else if val > root.Val {
		root.Right = root.Right.Insert(val)
	}
	return root
}

func (t *TreeNode) Remove(val int) *TreeNode {}
func (t *TreeSet) Remove(val int) *TreeSet   {}

// Time Complexity: O(logn)
// Space Complexity: O(h)
func (t *TreeSet) Search(val int) bool {
	return t.Root.Search(val)
}

// Time Complexity: O(logn)
// Space Complexity: O(h)
func (t *TreeSet) Insert(val int) {
	t.Root = t.Root.Insert(val)
}
