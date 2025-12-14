package main

// One Branch Rescursion
// Time Complexity: O(n)
// Space Complexity: O(n)
func factorial(n int) int {
	// Base case: n = 1
	if n == 1 {
		return 1
	} else {
		// Recursive case: n! = n * (n - 1)!
		return n * factorial(n-1)
	}
}
