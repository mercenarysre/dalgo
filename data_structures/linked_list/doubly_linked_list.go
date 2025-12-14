package main

import (
	"fmt"
)

type DoublyLinkedListNode struct {
	val  int
	next *DoublyLinkedListNode
	prev *DoublyLinkedListNode
}

type DoublyLinkedList struct {
	head *DoublyLinkedListNode
	tail *DoublyLinkedListNode
}

// instance of DoublyLinkedListNode
func NewDoublyLinkedListNode(val int) *DoublyLinkedListNode {
	return &DoublyLinkedListNode{
		val:  val,
		next: nil,
		prev: nil,
	}
}

// instance of DoublyLinkedList
func NewDoublyLinkedList() *DoublyLinkedList {
	// head and tail are dummy nodes in the linkedin list
	head := NewDoublyLinkedListNode(-1)
	tail := NewDoublyLinkedListNode(-1)
	head.next = tail
	tail.prev = head
	return &DoublyLinkedList{
		head: head,
		tail: tail,
	}
}

// Reading: accessing all nodes in the doubly linked list
func (d *DoublyLinkedList) Read() {
	if d.head == nil {
		fmt.Println("The doubly linked list is empty")
		return
	}

	curr := d.head
	for curr != nil {
		fmt.Println(curr.val)
		curr = curr.next
	}
}

// Searching: searching for a node value in the doubly linked list
func (d *DoublyLinkedList) Search(n int) {
	if d.head == nil {
		fmt.Println("The doubly linked list is empty")
		return
	}

	curr := d.head
	for curr != nil {
		if curr.val == n {
			fmt.Println("Found:", curr.val)
			return // stop after finding the value
		}
		curr = curr.next
	}
	fmt.Println("Not Found")
}

// Insertion in front of DoublyLinkedList
func (d *DoublyLinkedList) InsertFront(val int) {
	newNode := NewDoublyLinkedListNode(val)

	// case when the doubly linked list is empty
	if d.head == nil {
		fmt.Println("The doubly linked list is empty")
		d.head = newNode
		d.tail = newNode
		return
	}

	newNode.next = d.head
	d.head.prev = newNode
	d.head = newNode
}

// Insertion at end	of DoublyLinkedList
func (d *DoublyLinkedList) InsertEnd(val int) {
	newNode := NewDoublyLinkedListNode(val)

	if d.head == nil {
		d.head = newNode
		d.tail = newNode
		return
	}

	d.tail.next = newNode
	newNode.prev = d.tail
	d.tail = newNode
}

// Deletion	in front of DoublyLinkedList
func (d *DoublyLinkedList) RemoveFront() {
	if d.head == nil {
		fmt.Println("The doubly linked list is empty")
		return
	}

	if d.head == d.tail {
		d.head = nil
		d.tail = nil
		return
	}

	d.head = d.head.next
	d.head.prev = nil
}

// Deletion	at end	of DoublyLinkedList
func (d *DoublyLinkedList) RemoveEnd() {
	if d.head == nil {
		fmt.Println("The doubly linked list is empty")
		return
	}

	if d.head == d.tail {
		d.head = nil
		d.tail = nil
		return
	}

	d.tail = d.tail.prev
	d.tail.next = nil
}

func (d *DoublyLinkedList) Print() {
	// dummy head node, start printing from first element/node
	curr := d.head.next
	for curr != d.tail {
		fmt.Printf("%d -> ", curr.val)
		curr = curr.next
	}
	fmt.Println()
}
