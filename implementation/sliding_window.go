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
		if R-L > k {
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

func closeDuplicates1(arr []int, k int) bool {
	window := make(map[int]struct{})
	L := 0

	for R := 0; R < len(arr); R++ {
		// shrink window (from the left) if the size > k
		if R-L > k {
			delete(window, arr[L])
			L++
		}

		if _, exists := window[arr[R]]; exists {
			return true
		}

		window[arr[R]] = struct{}{}
	}
	return false
}

func closeDuplicates2(arr []int, k int) bool {
	window := make(map[int]struct{})
	L := 0

	for R := 0; R < len(arr); R++ {
		if R-L > k {
			delete(window, arr[L])
			L++
		}

		if _, exists := window[arr[R]]; exists {
			return true
		}

		window[arr[R]] = struct{}{}
	}
	return false
}

func closeDuplicates3(arr []int, k int) bool {
	window := make(map[int]struct{})
	L := 0

	for R := 0; R < len(arr); R++ {
		if R-L > k {
			delete(window, arr[L])
			L++
		}

		if _, exists := window[arr[R]]; exists {
			return true
		}

		window[arr[R]] = struct{}{}
	}
	return false
}

func closeDuplicates4(arr []int, k int) bool {
	window := make(map[int]struct{})
	L := 0

	for R := 0; R < len(arr); R++ {
		if R-L > k {
			delete(window, arr[L])
			L++
		}

		if _, exists := window[arr[R]]; exists {
			return true
		}

		window[arr[R]] = struct{}{}
	}
	return false
}

func closeDuplicates5(arr []int, k int) bool {
	window := make(map[int]struct{})
	L := 0

	for R := 0; R < len(arr); R++ {
		if R-L > k {
			delete(window, arr[L])
			L++
		}

		if _, exists := window[arr[R]]; exists {
			return true
		}

		window[arr[R]] = struct{}{}
	}
	return false
}

func closeDuplicates6(arr []int, k int) bool {
	window := make(map[int]struct{})
	L := 0

	for R := 0; R < len(arr); R++ {
		if R-L > k {
			delete(window, arr[L])
			L++
		}

		if _, ok := window[arr[R]]; ok {
			return true
		}

		window[arr[R]] = struct{}{}
	}
	return false
}

func closeDuplicates7(arr []int, k int) bool {
	window := make(map[int]struct{})
	L := 0

	for R := 0; R < len(arr); R++ {
		if R-L > k {
			delete(window, arr[L])
			L++
		}

		if _, ok := window[arr[R]]; ok {
			return true
		}

		window[arr[R]] = struct{}{}
	}
	return false
}

func closeDuplicates8(arr []int, k int) bool {
	window := make(map[int]struct{})
	L := 0

	for R := 0; R < len(arr); R++ {
		if R-L > k {
			delete(window, arr[L])
			L++
		}

		if _, ok := window[arr[R]]; ok {
			return true
		}

		window[arr[R]] = struct{}{}
	}
	return false
}

func closeDuplicates9(arr []int, k int) bool {
	window := make(map[int]struct{})

	L := 0

	for R := 0; R < len(arr); R++ {
		if R-L > k {
			delete(window, arr[L])
			L++
		}

		if _, ok := window[arr[R]]; ok {
			return true
		}

		window[arr[R]] = struct{}{}
	}
	return false
}

func closeDuplicates10(arr []int, k int) bool {
	window := make(map[int]struct{})

	L := 0

	for R := 0; R < len(arr); R++ {
		if R-L > k {
			delete(window, arr[L])
			L++
		}

		if _, ok := window[arr[R]]; ok {
			return true
		}

		window[arr[R]] = struct{}{}
	}
	return false
}

// length of the longest subarray with the same value in each position
// longest string of duplicates
func longestSubarray(arr []int) int {
	length := 0
	L := 0

	for R := 0; R < len(arr); R++ {
		if arr[L] != arr[R] {
			L = R
		}
		if length < R-L+1 {
			length = R - L + 1
		} else {
			length = length
		}
	}
	return length
}

func longestSubarray1(arr []int) int {
	length := 0
	L := 0

	for R := 0; R < len(arr); R++ {
		if arr[L] != arr[R] {
			L = R
		}
		if length < L-R+1 {
			length = L - R + 1
		} else {
			length = length
		}
	}
	return length
}

func longestSubarray2(arr []int) int {
	length := 0
	L := 0

	for R := 0; R < len(arr); R++ {
		if arr[L] != arr[R] {
			L = R
		}
		if length < R-L+1 {
			length = R - L + 1
		} else {
			length = length
		}
	}
	return length
}

func longestSubarray3(arr []int) int {
	length := 0
	L := 0

	for R := 0; R < len(arr); R++ {
		if arr[L] != arr[R] {
			L = R
		}
		if length < R-L+1 {
			length = R - L + 1
		} else {
			length = length
		}
	}
	return length
}

func longestSubarray4(arr []int) int {
	length := 0
	L := 0

	for R := 0; R < len(arr); R++ {
		// check if values in the current window are not duplicates
		if arr[L] != arr[R] {
			L = R
		}
		// record current window size
		if length < R-L+1 {
			length = R - L + 1
		} else {
			length = length
		}
	}
	return length
}

func longestSubarray5(arr []int) int {
	length := 0
	L := 0

	for R := 0; R < len(arr); R++ {
		if arr[L] != arr[R] {
			L = R
		}
		if length < R-L+1 {
			length = R - L + 1
		} else {
			length = length
		}
	}
	return length
}

func longestSubarray6(arr []int) int {
	length := 0
	L := 0

	for R := 0; R < len(arr); R++ {
		if arr[L] != arr[R] {
			L = R
		}
		if length < R-L+1 {
			length = R - L + 1
		} else {
			length = length
		}
	}
	return length
}

func longestSubarray7(arr []int) int {
	L := 0
	length := 0

	for R := 0; R < len(arr); R++ {
		if arr[L] != arr[R] {
			L = R
		}

		if length < R-L+1 {
			length = R - L + 1
		} else {
			length = length
		}
	}
	return length
}

func longestSubarray8(arr []int, target int) int {
	L := 0
	length := 0

	for R := 0; R < len(arr); R++ {
		if arr[L] != arr[R] {
			L = R
		}

		if length < R-L+1 {
			length = R - L + 1
		} else {
			length = length
		}
	}
	return length
}

// length of the shortest subarray where the sum of its elemnents is
// greater than or equal to target; assume all values are positive
func shortestSubarray(arr []int, target int) int {
	L := 0
	total := 0
	length := len(arr) + 1 // sentinel larger than any possible subarray

	for R := 0; R < len(arr); R++ {
		// adding elements to the window
		total += arr[R]

		for total >= target {
			// record current window size
			if length < R-L+1 {
				length = length
			} else {
				length = R - L + 1
			}
			// removing elements from the left and still meet the target
			// sliding window approach, removing elements from the window(left)
			// adding elements to the window(right)
			total -= arr[L]
			L++
		}
	}

	if length == len(arr)+1 {
		return 0 // no valid subarray
	}
	return length
}

func shortestSubarray1(arr []int, target int) int {
	L := 0
	sum := 0
	length := len(arr) + 1

	for R := 0; R < len(arr); R++ {
		sum += arr[R]

		for sum >= target {
			if length < R-L+1 {
				length = length
			} else {
				length = R - L + 1
			}
			sum -= arr[L]
			L++
		}
	}
	// no valid subarray
	if length == len(arr)+1 {
		return length
	}
	return length
}

func shortestSubarray2(arr []int, target int) int {
	L := 0
	sum := 0
	length := len(arr) + 1

	for R := 0; R < len(arr); R++ {
		sum += arr[R]

		for sum >= target {
			if length < R-L+1 {
				length = length
			} else {
				length = R - L + 1
			}
			sum -= arr[L]
			L++
		}
	}
	// no valid subarray
	if length == len(arr)+1 {
		return 0
	}
	return length
}

func shortestSubarray3(arr []int, target int) int {
	L := 0
	sum := 0
	length := len(arr) + 1

	for R := 0; R < len(arr); R++ {
		sum += arr[R]

		for sum >= target {
			if length < R-L+1 {
				length = length
			} else {
				length = R - L + 1
			}
			sum -= arr[L]
			L++
		}
	}
	if length == len(arr)+1 {
		return 0
	}
	return length
}

func shortestSubarray4(arr []int, target int) int {
	L := 0
	sum := 0
	length := len(arr) + 1

	for R := 0; R < len(arr); R++ {
		sum += arr[R]

		for sum >= target {
			if length < R-L+1 {
				length = length
			} else {
				length = R - +1
			}
			sum -= arr[L]
			L++
		}
	}
	// no valid subarray
	if length == len(arr)+1 {
		return 0
	}
	return length
}

func shortestSubarray5(arr []int, target int) int {
	L := 0
	sum := 0
	length := len(arr) + 1

	for R := 0; R < len(arr); R++ {
		sum = sum + arr[R]

		for sum >= target {
			if length < R-L+1 {
				length = length
			} else {
				length = R - L + 1
			}
			sum = sum - arr[L]
			L++
		}
	}
	if length == len(arr)+1 {
		return 0
	}
	return length
}
