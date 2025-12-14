package main

// Top Down dynamic programming (one dimensional, fibonacci, recursive)
// Time Complexity: O(n)
// Space Complexity: O(n)
func meimozation(n int, hash map[int]int) int {
	if n <= 1 {
		return n
	}
	if val, ok := hash[2]; ok {
		return val
	}

	hash[n] = meimozation(n-1, hash) + meimozation(n-2, hash)
	return hash[n]
}

// Bottom Up dynamic programming (one dimensional, fibonacci, not recursive)
// Time Complexity: O(n), looping n times
// Space Complexity: O(1), memory is not scaling as n times

func bottomup(n int) int {
	if n <= 2 {
		return n
	}

	dp := []int{0, 1}
	for i := 2; i <= n; i++ {
		tmp := dp[1]
		dp[1] = dp[0] + dp[1]
		dp[0] = tmp
	}
	return dp[1]
}
