// depth of a binary tree is the number of nodes along the
// longest path from the root node down to the farthest leaf node
package main

import "github.com/gammazero/deque"

// Time Complexity: O(n)
// Space Complexity: O(n)
func maxDepth(root *TreeNode) int {
	var queue deque.Deque
	if root == nil {
		return 0
	}
	queue.Pushback(root)
	level := 0

	for queue.Len() != 0 {
		levelLength := queue.Len()
		for i := 0; i < levelLength; i++ {
			curr := queue.PopFront().(*TreeNode)
			if curr.Left != nil {
				queue.Pushback(curr.Left)
			}
			if curr.Right != nil {
				queue.Pushback(curr.Right)
			}
		}
		level++
	}
	return level
}
