// i want to start by checking base cases, if root is nil, return false
// if subroot is nil, return false
// now there is a possiblity of the subRoot treenode being either
// on the left or right of the root treenode, i want to use the iterative DFS algorithm
// to deduce that--when root is equal to subroot
package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if subRoot == nil {
		return true
	}
	if root == nil {
		return false
	}

	if sameTree(root, subRoot) {
		return true
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func sameTree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}
	if root != nil && subRoot != nil && root.Val == subRoot.Val {
		return sameTree(root.Left, subRoot.Left) && sameTree(root.Right, subRoot.Right)
	}
	return false
}
