// edge case is if array length is less than or equal to 1
// i want to use a sliding window approach here,
// given the array is sorted, i want to start with two pointers
// a pointer starting at index 0 of the array
// and another pointer starting at index 1 of the array
// i want to compare the array values of both indexes
// if they are equal and if the difference between the
// pointers is greater than or equal to k, i want to append the array value to
// another array, then i want to increment the values of values of the left and right
// pointers respectively, also for values that are already added in the new array
// i want to ignore and not append
// the approach above is all about shifting from the left in the window
// and checking to the right adding to the window
// i return the new array

// Time Complexity: O(n)
// Space Complexity: O(n)
package main

import "sort"

func topKFrequent(nums []int, k int) []int {
	if len(nums) <= 1 {
		return nums
	}

	L := 0
	var nums1 []int
	nums2 := make(map[int]bool)

	for R := 1; R < len(nums); R++ {
		if nums[L] == nums[R] && R-L+1 >= k {
			if _, ok := nums2[nums[R]]; !ok {
				nums2[nums[R]] = true
				nums1 = append(nums1, nums[R])
			}
		}
		L++
	}
	return nums1
}

// Revised Solution
// Time Complexity:
// Space Complexity:
func topKFrequent1(nums []int, k int) []int {
	if len(nums) <= 1 {
		return nums
	}

	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}

	var arr [][2]int
	for num, cnt := range count {
		arr = append(arr, [2]int{cnt, num})
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] > arr[j][0]
	})

	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = arr[i][1]
	}

	return res
}
