// In Binary trees, a tree can be built from another data structure
// in O(nlogn) time	by inserting n elements into the tree
// Building a heap from another data stucture by inserting n elements
// at a faster pace in O(n) time and the heap built will satisfy
// the Strucuture property and Order property
// building a heap with n elements by pushing/inserting each element one-by-one
// using an efficient algorithm(heapify), perform the operation in O(n) time
func (h *Heap) Heapify(arr []int) {
	// 0-th position is moved to the end and is thus removed from the array
	arr = append(arr, arr[0])
	h.Heap = arr

	// cur determines the last
	cur := (len(h.Heap) - 1) / 2

	for cur > 0 {
		// Percolate Down
		i := cur
		for 2*i < len(h.Heap) {
			if 2*i+1 < len(h.Heap) && h.Heap[2*i+1] < h.Heap[2*i] && h.Heap[i] > h.Heap[2*i+1] {
				// Swap right child
				h.Heap[i], h.Heap[2*i+1] = h.Heap[2*i+1], h.Heap[i]
				i = 2*i + 1
			} else if h.Heap[i] > h.Heap[2*i] {
				// Swap left child
				h.Heap[i], h.Heap[2*i] = h.Heap[2*i], h.Heap[i]
				i = 2 * i
			} else {
				break
			}
		}
		cur--
	}
}