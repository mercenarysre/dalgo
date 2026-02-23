package main

// Time Complexity: O(n)
// Space Complexity: O(1)
func findMin(nums []int) int {
	minVal := nums[0]
	for _, num := range nums {
		if num < minVal {
			minVal = num
		}
	}
	return minVal
}
