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
