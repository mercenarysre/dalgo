// I want to start by defining base cases such that if either
// p or q is equal to nil, return false
// i want to intialize two arrays, then iterate through each binary
// tree, append each node val in the tree to an array
// iterate through one of the array using the for, range loop
// if the second array value is not equal to it, return false
// otherwise return true
package main

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return false
	}

	curr1 := p
	curr2 := q
	stack := []*TreeNode{}

	for curr1 != nil && curr2 != nil || len(stack) > 0 {
		if curr1 != nil && curr2 != nil {
			if curr1.Val != curr2.Val {
				return false
			}
			stack = append(stack, curr1)
			curr1 = curr1.Left
			curr2 = curr2.Left
		} else if curr1 == nil && curr2 == nil {
			curr1 = stack[len(stack)-1]
			curr2 = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			curr1 = curr1.Right
			curr2 = curr2.Right
		} else {
			return false
		}
	}
	return true
}
