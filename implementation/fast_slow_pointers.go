package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

type SinglyLinkedList struct {
	Head *ListNode
	Tail *ListNode
}

func NewSinglyLinkedList() *SinglyLinkedList {
	newnode := NewListNode(-1)
	return &SinglyLinkedList{
		Head: newnode,
		Tail: newnode,
	}
}

func (s *SinglyLinkedList) middleOfList() *ListNode {
	slow := s.Head
	fast := s.Head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func (s *SinglyLinkedList) hasCycle() bool {
	slow := s.Head
	fast := s.Head

	for fast != nil && fast.Next != nil {
		slow := slow.Next
		fast := fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

/ Determine if the linked list contains a cycle and
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
