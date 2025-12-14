package main

import (
	"fmt"
)

type ListNode struct {
	Val  int       // data
	Next *ListNode // next node: link to next node
}

// when dealing with a linked list, we only have immediate access to the list first node
// first node and last node are atrributes of the linked list class
type SinglyLinkedList struct {
	Head *ListNode
	Tail *ListNode
}

// instance of the Node returns the node's data and the next node in the list
//
//	All our program knows immediately is the memory address of the first node of the linked list
//	it doesn’t know offhand where any of the other nodes are.
func NewListNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

// linked list instance, a single node in the list: the head node and tail node will be the same
func NewSinglyLinkedList() *SinglyLinkedList {
	// new node instance
	node := NewListNode(-1)
	return &SinglyLinkedList{
		Head: node,
		Tail: node,
	}
}

// Linked List Reading: reading every node in the list
func (s *SinglyLinkedList) Read() {
	if s.Head == nil {
		fmt.Println("The singly linked list is empty")
		return
	}

	curr := s.Head
	for curr != nil {
		fmt.Println(curr.Val)
		curr = curr.Next
	}
}

// Linked List Searching: searching for a node value
func (s *SinglyLinkedList) Search(n int) {
	if s.Head == nil {
		fmt.Println("The singly linked list is empty")
		return
	}

	curr := s.Head
	for curr != nil {
		if curr.Val == n {
			fmt.Println("Found:", curr.Val)
			return // stop after finding the value
		}
		curr = curr.Next
	}
	fmt.Println("Not Found")
}

// Insertion at Beginning
func (s *SinglyLinkedList) InsertBeginning(val int) {
	newnode := NewListNode(val)
	if s.Head == nil {
		// list is empty, head and tail should point to the same node
		s.Head = newnode
		s.Tail = newnode
	} else {
		newnode.Next = s.Head
		s.Head = newnode
	}
}

// Insertion at End
func (s *SinglyLinkedList) InsertEnd(val int) {
	newnode := NewListNode(val)
	if s.Tail == nil {
		// list is empty, head and tail should point to the same node
		s.Head = newnode
		s.Tail = newnode
	} else {
		s.Tail.Next = newnode
		s.Tail = newnode
	}
}

// Deletion at Beginning
func (s *SinglyLinkedList) RemoveBeginning() {
	// case of where list is empty, head is nil
	if s.Head == nil {
		fmt.Println("The linked list is empty")
	}

	if s.Head == s.Tail {
		s.Head = nil
		s.Tail = nil
		return
	}

	s.Head = s.Head.Next
}

// Deletion at End
func (s *SinglyLinkedList) RemoveEnd() {
	if s.Head == nil {
		fmt.Println("The linked list is empty")
		return
	}

	if s.Head == s.Tail {
		s.Head = nil
		s.Tail = nil
		return
	}

	curr := s.Head
	for {
		if curr.Next == s.Tail {
			curr.Next = nil
			s.Tail = curr
			return
		}
		curr = curr.Next
	}
}

func (s *SinglyLinkedList) Print() {
	curr := s.Head.Next
	for curr != nil {
		fmt.Printf("%d -> ", curr.Val)
		curr = curr.Next
	}
	fmt.Println()
}
