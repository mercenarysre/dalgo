package main

import (
	"container/heap"
)

// -------- MinHeap --------
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] } // normal min-heap
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}

// -------- MaxHeap --------
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } // reverse → max-heap
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}

// -------- MedianFinder --------
type Median struct {
	small *MaxHeap // max-heap
	large *MinHeap // min-heap
}

func NewMedian() *Median {
	s := &MaxHeap{}
	l := &MinHeap{}
	heap.Init(s)
	heap.Init(l)
	return &Median{small: s, large: l}
}

// Time Complexity: O(logn)
// Space Complexity: O(n), number of elements in the stream
func (m *Median) Insert(num int) {
	// Push into max-heap (small)
	heap.Push(m.small, num)

	// Balance top elements: ensure all small <= all large
	// all values in the small(maxHeap) should be less than all
	// values in the large(minHeap), taking the max value in the small(maxHeap)
	// and make sure the value is <= the smallest value in the large(minHeap)
	if m.small.Len() > 0 && m.large.Len() > 0 && (*m.small)[0] > (*m.large)[0] {
		val := heap.Pop(m.small).(int)
		heap.Push(m.large, val)
	}

	// Handle uneven sizes
	// check the length of both heaps are equal or don't differ by one
	// else pop largest values from small(maxHeap) to large(minHeap)
	// or take smallest values from large(minHeap) to small(maxHeap)
	// in O(logn) time
	if m.small.Len() > m.large.Len()+1 {
		val := heap.Pop(m.small).(int)
		heap.Push(m.large, val)
	} else if m.large.Len() > m.small.Len()+1 {
		val := heap.Pop(m.large).(int)
		heap.Push(m.small, val)
	}
}

// Time Complexity: O(1)
// Space Complexity: O(1)
func (m *Median) GetMedian() float64 {
	// Calculate the median by checking the
	// length of both heaps if it is odd or even
	// if the values of both heaps are sorted
	// and find median
	if m.small.Len() > m.large.Len() {
		return float64((*m.small)[0])
	} else if m.large.Len() > m.small.Len() {
		return float64((*m.large)[0])
	}
	// Even count → average of tops
	return (float64((*m.small)[0]) + float64((*m.large)[0])) / 2
}
