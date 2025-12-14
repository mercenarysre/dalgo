package main

// Time Complexity: O(logn)
// Space Complexity: O(1)
// Binary Search for Ordered Array
func binarySearch(arr []int, target int) int {
	// lower bound
	L := 0
	// upper bound
	R := len(arr) - 1

	var mid int

	// loop runs when we have a range of elements where the target may lie,
	// the algorithm keeps narrowing that range,
	// inspecting the middlemost value of upper/lower bounds
	// The clause L <= R will be reached when there’s no
	// more range left, and we can conclude that the target is not present in the array.
	for L <= R {
		// finding the midpoint, result of division of integers is rounded up to the nearest integer
		mid = (L + R) / 2

		// inspecting the value at midpoint == target
		// if not, changing values of upper/lower bounds
		if target > arr[mid] {
			L = mid + 1
		} else if target < arr[mid] {
			R = mid - 1
		} else {
			return mid
		}
	}
	// having narrowed down the bounds until they have reached each other,
	// that means the value we are searching is not in the array
	return -1
}

func binarySearch1(arr []int, target int) int {
	// left pointer
	L := 0

	// right pointer
	R := len(arr) - 1

	var mid int

	for L <= R {
		mid = (L + R) / 2

		if target < arr[mid] {
			R = mid - 1
		} else if target > arr[mid] {
			L = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

// Time Complexity: O(logn)
// Space Complexity: O(1)
// Binary Search variation for a range of numbers rather than an array,
// without a specific target
func binarySearchRange(low, high int) int {
	var mid int

	for low <= high {
		mid = (low + high) / 2
		switch {
		case isCorrect(mid) > 0:
			high = mid - 1
		case isCorrect(mid) < 0:
			low = mid + 1
		default:
			return mid
		}
	}
	return -1
}

func isCorrect(n int) int {
	switch {
	case n > 10:
		return 1
	case n < 10:
		return -1
	default:
		return 0
	}
}
