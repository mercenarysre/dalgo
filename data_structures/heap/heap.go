package main

// heaps are implemeted using arrays uder the hood
type Heap struct {
	Heap []int
}

// constructor function
func NewHeap() *Heap {
	return &Heap{
		Heap: append([]int{}, 0),
	}
}

// focusing on Binary Heap, a specific kind of binary tree, Max Heap/Min Heap
// Heap from a Priority Queue perspective has a primary function to allow access
// to the highest priority item in the queue(root node), each time we take care of
// the highest priority item and remove it from the loop, the next highest item
// floats to the top of the heap and is on deck to be addressed next.

// Reading: looking at the value of the root node/last node
// With heaps we can read the minimum or maximum value (highest priority) in constant time,
//	O(1)by simply reading the element at the first position,
func (h *Heap) Top() int {
	if len(h.Heap) > 0 {
		return h.Heap[0]
	}
	return -1
}

// Searching: require us to inspect each node, so search is
// not an operation usually implemented in the context of heaps.

// Insertion and Deletion are so dependent on the last node
// To solve the problem of the last node
// Initialize the heap as an empty array
// A root-node method which returns the first node/item of the array
// A last-node method which returns the last node/item of the array
// Traversing an Array-Based-Heap, enables traversing the Heap without doing O(N) steps
// after assigning indexes to the heap nodes with root node	starting from 0 and then going breadth-wise
// node left child: (index * 2) + 1
// node right child: (index * 2) + 2
// node parent: (index - 1) / 2
// These formulas only work when the tree is a complete binary tree and the
// array is filled contiguously from left to right

// The last node is critical to the heap primary operations, the computer can't see it
// and we have to make sure that finding the last node is efficient, thus heaps are implemented as arrays

// Insertions:  Turn value into last node by inserting it at the end of the arrayand taking subsequent steps
// Min Heap Insertion
// Time Complexity: O(logn) // it travels from the leaf to the root, worst case, height of heap = logn
// Space Complexity: O(1)
func (h *Heap) Push(val int) {
	h.Heap = append(h.Heap, val)
	i := len(h.Heap) - 1

	// Percolate up
	for i > 0 && h.Heap[i] < h.Heap[(i-1)/2] {
		// Swap
		h.Heap[i], h.Heap[(i-1)/2] = h.Heap[(i-1)/2], h.Heap[i]
		i = (i - 1) / 2
	}
}

// Max	Heap Insertion
func (h *Heap) Push(val int) {
	h.Heap = append(h.Heap, val)
	i := len(h.Heap) - 1

	// Percolate up
	for i > 0 && h.Heap[i] > h.Heap[(i-1)/2] {
		// Swap
		h.Heap[i], h.Heap[(i-1)/2] = h.Heap[(i-1)/2], h.Heap[i]
		i = (i - 1) / 2
	}
}

// Deletions, only ever delete the root node and do subsequent steps
// Min	Heap Deletion
// Time Complexity: O(logn), percolating down — the element may move from the root all the way down to a leaf
// height = O(log n)
// Space Complexity: O(1)
func (h *Heap) Pop() int {
	if len(h.Heap) == 0 {
		return -1
	}

	if len(h.Heap) == 1 {
		res := h.Heap[0]
		h.Heap = h.Heap[:0]
		return res
	}

	res := h.Heap[0]
	h.Heap[0] = h.Heap[len(h.Heap)-1]
	h.Heap = h.Heap[:len(h.Heap)-1]

	i := 0
	// Percolate down
	for 2*i+1 < len(h.Heap) {
		if 2*i+2 < len(h.Heap) && h.Heap[2*i+2] < h.Heap[2*i+1] && h.Heap[i] > h.Heap[2*i+2] {
			// Swap right child
			h.Heap[i], h.Heap[2*i+2] = h.Heap[2*i+2], h.Heap[i]
			i = 2*i + 2
		} else if h.Heap[i] > h.Heap[2*i+1] {
			// Swap left child
			h.Heap[i], h.Heap[2*i+1] = h.Heap[2*i+1], h.Heap[i]
			i = 2*i + 1
		} else {
			break
		}
	}
	return res
}

// Deletion, only ever delete the root node and do subsequent steps
// Max Heap Deletion
// Time Complexity: O(logn)
// Space Complexity: O(1)
func (h *Heap) Pop() int {
	if len(h.Heap) == 0 {
		return -1
	}

	if len(h.Heap) == 1 {
		res := h.Heap[0]
		h.Heap = h.Heap[:0]
		return res
	}

	res := h.Heap[0]
	h.Heap[0] = h.Heap[len(h.Heap)-1]
	h.Heap = h.Heap[:len(h.Heap)-1]

	i := 0
	// Percolate down
	for 2*i+1 < len(h.Heap) {
		if 2*i+2 < len(h.Heap) && h.Heap[2*i+2] > h.Heap[2*i+1] && h.Heap[i] < h.Heap[2*i+2] {
			// Swap right child
			h.Heap[i], h.Heap[2*i+2] = h.Heap[2*i+2], h.Heap[i]
			i = 2*i + 2
		} else if h.Heap[i] < h.Heap[2*i+1] {
			// Swap left child
			h.Heap[i], h.Heap[2*i+1] = h.Heap[2*i+1], h.Heap[i]
			i = 2*i + 1
		} else {
			break
		}
	}
	return res
}
