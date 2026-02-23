// i want to approach this problem by first sorting the array of
// integers, so that i can easily get the consecutive sequence
// of elements exactly 1 greater than each other
// then to actually get the consecutive sequence, i want to create an hashmap
// i want to start by iterating the sorted array and use a condition whereby if
// next integer in the array is greater than the current integer
// by 1, i want to add the element(integer)	to an hashmap, if the integer is
// already in the hashmap, i don't want to put it again
// then at the end of the iterations, i want to return the length of the hashmap

// Time Complexity: O(n2)
// Space Complexity: O(k)
package main

func longestConsecutive(nums []int) int {
	for i := 1; i < len(nums); i++ {
		j := i - 1

		for j >= 0 && nums[j+1] < nums[j] {
			tmp := nums[j+1]
			nums[j+1] = nums[j]
			nums[j] = tmp
			j--
		}
	}

	window := make(map[int]struct{})
	n := len(nums) - 1
	for i := 0; i < n; i++ {
		j := i + 1
		if nums[i]+1 == nums[j] {
			if _, ok := window[nums[i]]; !ok {
				window[nums[i]] = struct{}{}
			}
			if _, ok := window[nums[j]]; !ok {
				window[nums[j]] = struct{}{}
			}
		}
	}
	return len(window)
}
