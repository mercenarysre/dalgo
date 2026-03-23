// an edge case i want to check is if the head	of the linked list
// is equal to nil; if yes, i want to return head which is nil
// another edge case is if the the head of the linked list is
// equal to the tail of the linked list then i want to set the
// head and the tail to be equal to nil, and return either head or tail
// i want to start bv creating a list(array) of linked lists, and then
// appending each linked list to the array, after that the problem lies
// in identifying the index of the linked list we want to remove, the logic
// to doing this is to subtract the given n index from the length of the
// linked list array, after that we determine if the subtracted index is 0--
// then return the next node, otherwise if the subtracted index is not
// 0, i want to check the node before that subtracted index and say its Next pointer
// points to the node after the subtracted index node
package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// Time Complexity: O(n)
// Space Complexity: O(n)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	nodes := []*ListNode{}
	cur := head
	for cur != nil {
		nodes = append(nodes, cur)
		cur = cur.Next
	}

	removeIndex := len(nodes) - n
	if removeIndex == 0 {
		return head.Next
	}

	nodes[removeIndex-1].Next = nodes[removeIndex].Next
	return head
}
