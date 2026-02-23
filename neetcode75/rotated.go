package main

// Time Complexity: O(logn)
// Space Complexity: O(1)
func search(nums []int, target int) int {
	L := 0
	R := len(nums) - 1
	var M int

	for L <= R {
		M = (L + R) / 2
		if nums[M] == target {
			return M
		}

		if nums[L] <= nums[M] {
			if target >= nums[L] && target < nums[M] {
				R = M - 1
			} else {
				L = M + 1
			}
		} else {
			if target > nums[M] && target <= nums[R] {
				L = M + 1
			} else {
				R = M - 1
			}
		}
	}
	return -1
}
