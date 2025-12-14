package main

// Variable sliding window; when we don't have a fixed window size
// and we need to keep expanding the window as long as the
// window meets a certain constraint

// Find the length of the longest subarray wih the
// same value in each position; longest string of duplicates
// Time Complexity: O(n)
// Space Complexity: O(1)
func longestSubarray(arr []int) int {
	length := 0
	L := 0

	for R := 0; R < len(arr); R++ {
		if arr[L] != arr[R] {
			L = R
		}
		length = max(length, R-L+1)
	}
	return length
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// Find the minimum length subarray where the sum is greater than or equal to the target.
// Assume all values in the input are positive.
// Time Complexity: O(n), even though we have nested loops, the inner loop won't neccessarily
// run n times at every iteration==amortized analysis
// Space Complexity: O(1)
func shortestSubarray(arr []int, target int) int {
	L := 0
	total := 0
	length := len(arr) + 1 // sentinel larger than any possible subarray

	for R := 0; R < len(arr); R++ {
		total += arr[R]

		for total >= target {
			length = min(length, R-L+1)
			total -= arr[L]
			L++
		}
	}

	if length == len(arr)+1 {
		return 0 // no valid subarray
	}
	return length
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
