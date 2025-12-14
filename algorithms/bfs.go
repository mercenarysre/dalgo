package main

import (
	"fmt"

	"github.com/gammazero/deque"
)

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

// For breath-first search (BFS), we prioritize breadth,
// meaning we focus on visiting all the nodes on one level before moving on to the next level
// BFS is also known as level-order traversal when referring to trees, since we visit the nodes level by level

// A queue data structure, more specifically, a deque,
// allows us to remove elements both from the head and the tail in O(1) time.
// For BFS we will append elements to the tail and remove elements from the head
// as we go through each level of the tree from left to right.

// Time Complexity: O(n), n is the number of nodes in the tree,
// we visit every node exactly once.
// Space Complexity: O(n), n is the number of nodes in the tree,
// We store an entire level of the tree in the queue at a time.
// In the worst case the last level may be roughly half the size of the tree, so the space complexity is O(n).

// For a balanced tree, maximum size of the queue happens when BFS reaches the widest level of the tree
// The last level (deepest level) can have up to n/2 nodes in a full/balanced binary tree.
// Because in a full tree, about half the nodes live at the bottom
// In a perfect or full balanced binary tree, the last level can contain
// about half the total nodes of the tree (n/2).
//
//	For example, in a tree with 15 nodes:
//	Last level = 8 nodes (which is ≈ 15/2).
//	But in Big-O notation, we only care about the order of growth, not constants.
//	n/2 grows in the same way as n,	so O(n/2) → O(n).

func (root *TreeNode) BFS() {
	// Initially, we append the root node to our queue.
	// We then enter a while loop that runs as long as our queue is not empty.
	// We print the level we are currently on.
	// We loop through the queue and remove nodes in the current level.
	// If the node has children, we append them to the queue from left to right.
	// After the current level is done, we increment the level by 1.
	// Our queue becomes empty once we have visited all of the nodes and the outer while loop will terminate.
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
