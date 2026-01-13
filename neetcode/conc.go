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
