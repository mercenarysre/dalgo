package main

// Fast and Slow pointers algorithm technique
// applicable to linked lists

type ListNode struct {
	Val  int       // data
	Next *ListNode // next node: link to next node
}

type SinglyLinkedList struct {
	Head *ListNode
	Tail *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

func NewSinglyLinkedList() *SinglyLinkedList {
	// new node instance
	node := NewListNode(-1)
	return &SinglyLinkedList{
		Head: node,
		Tail: node,
	}
}

// Find the middle of a linked list with two pointers.
// Time Complexity: O(n)
// Space Complexity: O(1)
func (s *SinglyLinkedList) middleOfList() *ListNode {
	slow := s.Head
	fast := s.Head

	// fast pointer is not null
	// fast.next pointer is not null
	// when fast.next pointer is null, we stop the loop
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// Determine if the linked list contains a cycle.
// Time Complexity: O(n), slow pointer will move O(n), fast pointer will move O(2.n) still goes to O(n)
// Space Complexity: O(1)
func (s *SinglyLinkedList) hasCycle() bool {
	slow := s.Head
	fast := s.Head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// Determine if the linked list contains a cycle and
// return the beginning of the cycle, otherwise return null.
// Time Complexity: O(n)
// Space Complexity: O(1)
func (s *SinglyLinkedList) headCycle() *ListNode {
	slow := s.Head
	fast := s.Head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}

	if fast == nil && fast.Next == nil {
		return nil
	}

	slow2 := s.Head
	for slow != slow2 {
		slow = slow.Next
		slow2 = slow.Next
	}
	return slow
}
