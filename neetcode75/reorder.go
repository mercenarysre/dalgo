package main

/*
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// Time Complexity: O(n)
// Space Complexity: O(n)
func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	nodes := []*ListNode{}
	cur := head
	for cur != nil {
		nodes = append(nodes, cur)
		cur = cur.Next
	}

	L := 0
	R := len(nodes) - 1
	for L < R {
		nodes[L].Next = nodes[R]
		L++
		if L >= R {
			break
		}
		nodes[R].Next = nodes[L]
		R--
	}

	nodes[L].Next = nil
}
