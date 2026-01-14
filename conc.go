// I got the length of the original array
// I initialized a new array with a length 2x of the original array
// I iterate through the length of the first array using a for loop
// assign the value(s) of the first array to the new array
// simulataneuosly assigning the value(s) of the first array
// to the additional length of the new array
//return the new array

// Time Complexity: O(n)
// Space Complexity: O(n)
package main

func getConcatenation(nums []int) []int {
	n := len(nums)
	nums1 := make([]int, 2*n)
	for i := 0; i < n; i++ {
		nums1[i] = nums[i]
		nums1[i+n] = nums[i]
	}
	return nums1
}
