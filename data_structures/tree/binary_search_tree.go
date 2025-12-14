package main

import (
	"fmt"

	"github.com/gammazero/deque"
)

// A binary search tree is a tree:
// where a (root) node has at least one left child and one last child
// where a (root) node left descandants can only contain value(s) that are less than the node (value) itself
// and the node right descandants can only contain value(s) that are greater than the node (value) itself
// binary search tree is a variation of a binary tree with the addition of a sorted property
// Asides search which is O(logn) BSTs are preferred over sorted arrays because they allow
// for insertion and deletion in O(logn), (This assumes that the tree is balanced)
// which are o(n) operations in sorted arrays.
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

// Reading: Binary Search Tree Traversal
// Time Complexity: O(n), visiting every single node(the size of the tree)
// regardless of the height of the tree
// Space Complexity: O(h), height of the tree which would be O(logn) for a balanced binary tree
// or O(n) for a skewed

// Searching: Three methods of Searching: the method below, DFS or BFS.
// Time Complexity: O(logn) if the tree is balanced
// Space Complexity: O(h), where h is the height of the tree,
func (root *TreeNode) Search(target int) bool {
	if root == nil {
		return false
	}

	// In a binary search tree as you search each node you enter
	// becomes a root, right or left node of its respective sub-tree
	if target > root.Val {
		return root.Right.Search(target)
	} else if target < root.Val {
		return root.Left.Search(target)
	}
	return true
}

// Depth First Search: PreOrder, InOrder, PostOrder
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
// Time Complexity: O(n), we visit each node exactly once
// Space Complexity: O(n), we implement a queue, we store an entire level of the tree in the queue at a time
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

// Insertion
// Time Complexity: O(logn) if the tree is balanced, O(n) if the tree is unbalanced
// Space Complexity: O(logn) in the best case if the tree is balanced;
// O(n) in the worst case, if the tree is unbalanced--this is because
// we are using recursion to traverse the tree, and
// the number of recursive calls is proportional to the height of the tree.
func (root *TreeNode) Insert(val int) *TreeNode {
	// if the current node is null, we return a new node with the value val
	if root == nil {
		return NewTreeNode(val)
	}

	if val > root.Val {
		root.Right = root.Right.Insert(val)
	} else if val < root.Val {
		root.Left = root.Left.Insert(val)
	}
	return root
}

// Deletion: a recursive function, deletion with no child, one child, two children

//	If the node being deleted has no children, simply delete it.

//	If the node being deleted has one child, delete the node and plug the child
//	into the spot where the deleted node was.

//	When deleting a node with two children, replace the deleted node with
//	the successor node. The successor node is the child node whose value is
//	the least of all values that are greater than the deleted node.
//	To find the successor node: visit the right child of the deleted value,
// 	if the right child of the value deleted has a left child (and)
//	then keep on visiting the left child of each subsequent child until there
// 	are no more left children. The bottom value is the successor node.
//	If the successor node has a right child, after plugging the successor node
// 	into the spot of the deleted node, take the former right child of the
// 	successor node and turn it into the left child of the former parent of the successor node
// 	Remove a node and return the root of the BST.

// Time Complexity: O(logn) if the tree is balanced, O(n) if the tree is unbalanced
// Space Complexity: O(logn) in the best case if the tree is balanced;
// O(n) in the worst case, if the tree is unbalanced
// this is because we are using recursion to traverse the tree, and
// the number of recursive calls is proportional to the height of the tree.
// REWRITE THIS CODE
func (root *TreeNode) Remove(val int) *TreeNode {
	if root == nil {
		return nil
	}
	if val > root.Val {
		root.Right = root.Right.Remove(val)
	} else if val < root.Val {
		root.Left = root.Left.Remove(val)
	} else { // val == root.Val
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			minNode := root.Right.MinValueNode()
			root.Val = minNode.Val
			root.Right = root.Right.Remove(minNode.Val)
		}
	}
	return root
}

// Return the minimum value node of the BST.
func (root *TreeNode) MinValueNode() *TreeNode {
	curr := root
	for curr != nil && curr.Left != nil {
		curr = curr.Left
	}
	return curr
}

// Return the maximum value node of the BST.
func (root *TreeNode) MaxValueNode() *TreeNode {
	curr := root
	for curr != nil && curr.Right != nil {
		curr = curr.Right
	}
	return curr
}
