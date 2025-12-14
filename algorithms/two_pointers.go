package main

// For a sliding window we care about the entire window
// and values in the window

// Two pointer algorithms focuses on two pointers, a left pointer
// and a right pointer, both starting at some indexes of an array
// They don't always have to start at the beginning of the array,
// they can start at any index of the array, increment/decrement L/R based on the given problem

// Given a string of characters, return true if it's a palindrome,
// Time Complexity: O(n)
// Space Complexity: O(1)
func isPalindrome(arr []int) bool {
	L := 0
	R := len(arr) - 1

	for L < R {
		if arr[L] != arr[R] {
			return false
		}
		L += 1
		R -= 1
	}
	return true
}

// Given a sorted array of integers, return the indices
// of two elements (in different positions) that sum up to
// the target value. Assume there is exactly one solution.
// Time Complexity: O(n)
// Space Complexity: O(1)
// Assumes array is sorted
func targetSum(arr []int, target int) (int, int) {
	L := 0
	R := len(arr) - 1

	for L < R {
		if arr[L]+arr[R] > target {
			R -= 1
		} else if arr[L]+arr[R] < target {
			L += 1
		} else {
			return L, R
		}
	}
	return -1, -1
}
