// an edge case is when the linked list (head) is empty, return nil
// i want to loop over the linked list such that i want to run three pointers
// a curr pointer which is not empty and is the current node we are processing
// a prev pointer which is the node that should come after curr once reversed
// a temp pointer which is the original next node
// the cycle will be such that the first node Next pointer is pointed to nil, the
// second node Next pointer will be pointed to the first node, the third node Next
// pointer will be pointed to the second node and so on...when the curr node is nil
// the cycle stops and we return the prev pointer which contains the reversed list
package main

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var prev *ListNode
	curr := head
	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}
	return prev
}
