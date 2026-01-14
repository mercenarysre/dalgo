// create an hashmap, hashmap of key type int and value type bool
// iterate through array elements, if the the array element already exist in the hashmap
// i want to return the element hash value
// if the array element does not exists in the hashmap
// i want to add the element to the hashmap
// if there are no elements that appear more than once
// i want to return false

package main

func hasDuplicate(nums []int) bool {
	nums1 := make(map[int]bool)

	for _, num := range nums {
		_, ok := nums1[num]
		if ok {
			return true
		}

		if !ok {
			nums1[num] = true
		}
	}
	return false
}
