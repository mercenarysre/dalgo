package main

// Given an array, return true if there are two elements
// within a window of size k that are equal
// Time Complexity: O(n) single pass on the array and our hashset allows us to have O(1) lookup
// Space Complexity: O(k) because we are storing at most k distinct elements in our hashset.
func closeDuplicates(arr []int, k int) bool {

	// acts as a hashet
	window := make(map[int]struct{})

	// left pointer
	L := 0

	// right pointer
	for R := 0; R < len(arr); R++ {

		// shrink window if size > k
		if R-L+1 > k {
			// remove element that is no longer in the window and shift left pointer of the array
			delete(window, arr[L])
			L++
		}

		// check duplicate
		if _, exists := window[arr[R]]; exists {
			return true
		}

		// add current element
		window[arr[R]] = struct{}{}
	}
	return false
}
