// an edge case if it the length of the array is less than three
// i return an empty two dimensional array

// to begin with, a brute force approach is iterate through the array
// using a loop by picking the first element in the array, then a second loop
// that uses the second element in the array, then a third loop that uses
// the third element in the array; if the three elements are equal to zero
// i want to put the elements in a subarray, which i will append to a hashmap
// if such three elements exists in the loop, i want to confirm it from the hashmap
// if its doesn't i want to append it to the hashmap in a subarray else i want to ignore
// and continue
// then i want to create a two dimensional array of which i will iterate through the keys/
// values of the hashmap and append the values to the two dimensional array, and return the array

// Time Complexity: O(n3), n is the length of the given array
// Space Complexity: O(m), where m is the number of triplets
package main

func threeSum(nums []int) [][]int {
	nums1 := map[[3]int]struct{}{}

	// sorting to handle duplicates so i don't have to check duplicates against hashmap later
	// and also to only ever generate one ordering per triplet
	for i := 0; i < len(nums); i++ {
		j := i - 1
		for j >= 0 && nums[j+1] < nums[j] {
			tmp := nums[j+1]
			nums[j+1] = nums[j]
			nums[j] = tmp
			j--
		}
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					nums1[[3]int{nums[i], nums[j], nums[k]}] = struct{}{}
				}
			}
		}
	}

	var nums2 [][]int
	for k, _ := range nums1 {
		nums2 = append(nums2, []int{k[0], k[1], k[2]})
	}
	return nums2
}
