// the analogy here is to use two pointers for this problem
// whereby a slow pointer is assigned to the list head, and
// a fast pointer is assigned to the list head, now we start
// a loop to iterate the list such that the fast pointer is a
// node ahead of the slow pointer, the fast pointer is not nil
// and its next pointer is not nil as well, a cycle occurs when
// the slow and fast pointer meets at a node, and in this case
// this means the tail node connects to the 1st node, however if the
// two pointers does not meet given the loops ends--there is no cycle
// and return false

// Time Complexity: O(n)
// Space Complexity: O(1)
package main

func hasCycle(head *ListNode) bool {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
