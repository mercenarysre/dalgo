// given array is sorted and DISTINCT
// an edge case is if the array length is equal to 0, there is no target
// i want to declare left and right pointers for the array, the left pointer
// is is the first index of the array, the right pointer is the last index
// of the array; now i want do a binary search on the array such that i want
// to declare a variable M which is the half(mid-value) of the left pointer and the right point
// searching for the target implies i want if the mid value is equal to the target
// then return the mid value, if the target is greater than the mid value it means
// the target lies on the other(higher) half on the mid value, the left pointer becomes the mid value
// + 1, otherwise if the target is less than the mid value it the means the target lies on the
// other(lower) half on the mid value, after iterating through all the array integers and did not
// find the target, return -1

package main

func search(nums []int, target int) int {
	L := 0
	R := len(nums) - 1

	var M int

	for L <= R {
		M = (L + R) / 2

		if target > nums[M] {
			L = M + 1
		} else if target < nums[M] {
			R = M - 1
		} else {
			return M
		}
	}
	return -1
}
