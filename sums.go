// edge case: if initial array length is equal to 1, return nil

// i want to iterate through the array using a for loop, take the first element
// using another for loop, i want to iterate the other elements of the array
// then check if the sum of the first loop element and the second loop element(s)
// is equal to the target, get the indices of both elements; create an array for the
// indices and append the indices to the array, (there is a caveat however such that
// the first loop only reaches the second to the last element of the array and the
// second loop reaches the end the array) return the array, otherwise return nil
// exactly one pair
// return answer with the smaller index first

// Time Complexity: O(n2)
// Space Complexity: O(1), data structure with constant space, storing only two elements, not increasing in size
package main

func twoSum(nums []int, target int) []int {
	if len(nums) == 0 || len(nums) == 1 {
		return nil
	}

	var nums1 []int
	n := len(nums) - 2

	for i := 0; i <= n; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				nums1 = append(nums1, i, j)
				return nums1
			}
		}
	}
	return nil
}
