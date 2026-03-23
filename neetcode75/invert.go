// This problem employs the breadth first search
// approach such that we want to append the root node
// to a queue, we then enter a for	loop that runs
// as long as the queue is not empty, we start another loopn
// that iterates from zerp to the last elememt of the queue;
// we pop an element from the queue and we print its value
// this is such that we are printing the element valueof the
// of the current level, also we do this such that when a node has children, we append
// them to the queue from right to left, after the current level
// is done, we increment the level by 1, the queue becomes
// empty once we have visited all the nodes
// and the outer(first) loop will terminate

// Time Complexity: O(n)
// Space Complextity: O(n)
package main

import (
	"fmt"

	"github.com/gammazero/deque"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
	var queue deque.Deque
	if root != nil {
		queue.Pushback(root)
	}

	level := 0
	for queue.Len() != 0 {
		levelLength := queue.Len()
		for i := 0; i < levelLength; i++ {
			curr := queue.PopFront()
			fmt.Println(curr.Val)
			if curr.Right != nil {
				queue.Pushback(curr.Right)
			}
			if curr.Left != nil {
				queue.Pushback(curr.Left)
			}
		}
		level++
	}
	return root
}

// Alternative Solution
func invertTree(root *TreeNode) *TreeNode {
	var queue deque.Deque
	if root != nil {
		queue.Pushback(root)
	}

	level := 0
	for queue.Len() != 0 {
		levelLength := queue.Len()
		for i := 0; i < levelLength; i++ {
			curr := queue.PopFront().(*TreeNode)
			curr.Left = curr.Right
			curr.Right = curr.Left
			if curr.Left != nil {
				queue.Pushback(curr.Left)
			}
			if curr.Right != nil {
				queue.Pushback(curr.Right)
			}
		}
		level++
	}
	return root
}
