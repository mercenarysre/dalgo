package main

//	Backtracking is an algorithm with a lot of overlap with DFS.
//	It operates on a brute-force approach which is to try all possible solutions and backtrack
// when we hit a dead-end.

//	Imagine that we are trapped in a maze and we are trying to find our way out.
// We can try all possible paths until we find the correct one.
// If we hit a dead-end, we backtrack and try another path. This is the essence of backtracking.

// Definition of TreeNode
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

//	determine if there exists a path from the root to a leaf node without having a value of 0 in the path.
// 	we cannot have a node with value 0 in our path.

// Time Complexity: O(n), because in the worst case we may need to visit all the nodes in the tree
// Space Complexity: O(h)	h is the height of the tree. This is because we are using recursion
// and the maximum depth of the recursion is the height of the tree.
func (root *TreeNode) CanReachLeaf() bool {

	// We also know that if the tree is empty or root value is 0, then there cannot exist a valid path either
	if root == nil || root.Val == 0 {
		return false
	}

	// if we reach a leaf node we can return true since it means there is a path that exists from root to leaf
	// reach a leaf	node children, the node itself is not nil and its value is not equal to zero
	if root.Left == nil && root.Right == nil {
		return true
	}

	//	If the left-subtree returns true, we can return true as well.
	//	If the left-subtree returns false, we can explore the right-subtree.
	//	If the right-subtree returns true, we can return true as well.
	//	If both the left and right subtrees return false, we can return false as well.

	if root.Left.CanReachLeaf() {
		return true
	}

	if root.Right.CanReachLeaf() {
		return true
	}

	return false
}

//	determine if there exists a path from the root to a leaf node without having a value of 0 in the path
// 	and building the path if it exists instead of just returning true or false
func (root *TreeNode) LeafPath(path *[]int) bool {
	if root == nil || root.Val == 0 {
		return false
	}
	//	storing the nodes along the path in the path list when true
	//  removing the nodes when false
	*path = append(*path, root.Val)

	if root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left.LeafPath(path) {
		return true
	}
	if root.Right.LeafPath(path) {
		return true
	}

	// remove the last element in path if both recursive calls return false
	*path = (*path)[:len(*path)-1]
	return false
}
