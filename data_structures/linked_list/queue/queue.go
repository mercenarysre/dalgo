package main

import (
	"fmt"
)

type Queue struct {
	queue []int // a queue is an array with three restrictions
}

// only the front element of a queue can be read
func (s *Queue) Read() int {
	if len(s.queue) == 0 {
		fmt.Println("Error: Queue is empty")
		return 0
	}
	front := s.queue[0]
	return front
}

// data can be inserted only at the end of a queue
func (s *Queue) Push(n int) {
	s.queue = append(s.queue, n)
}

// data can be deleted only at the front of a queue
func (s *Queue) Pop() int {
	if len(s.queue) == 0 {
		fmt.Println("Error: Queue is empty")
		return 0
	}
	// deleting an element from the front of an array, front of a queue
	front := s.queue[0]
	s.queue = s.queue[1:]

	return front
}

// size of the queue
func (s *Queue) Size() int {
	return len(s.queue)
}

/*
package main

import (
	"fmt"
	"os"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type Queue struct {
	Left  *ListNode // front of Queue
	Right *ListNode // back of Queue
}

// Instance of a ListNode
func NewListNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

// Instance of a Queue
func NewQueue() *Queue {
	// Left and Rights are pointers in the Queue, implying the Queue is empty
	return &Queue{
		Left:  nil,
		Right: nil,
	}
}

// read data at the front of a queue

// insert data at the end of a queue
func (q *Queue) Enqueue(val int) {
	newNode := NewListNode(val)
	if q.Right != nil {
		// Queue is not empty
		q.Right.Next = newNode
		q.Right = q.Right.Next // newNode
	} else {
		// Queue is empty
		q.Left = newNode
		q.Right = newNode
	}
}

// delete data from the front of a queue
func (q *Queue) Dequeue() int {
	if q.Left == nil {
		os.Exit(0)
	}
	val := q.Left.Val
	q.Left = q.Left.Next
	if q.Left == nil {
		q.Right = nil
	}
	return val
}

func (q *Queue) Print() {
	curr := q.Left
	for curr != nil {
		fmt.Printf("%d -> ", curr.Val)
		curr = curr.Next
	}
	fmt.Println()
}
*/
