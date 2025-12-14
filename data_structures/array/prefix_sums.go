package main

// Prefix is a contiguous subarray that starts at the beginning of an array
// Prefix sums tells the sum of contiguous subarrays that
// happen at the beginning at an array; not the sum of subarrays
// that does not start from the beginning of the array
// Prefix Sum
// Prefix Product
// Postfix Sum
// Postfix Product

// Given an array of values, design a data structure that can query
// the sum of the subarrarys of the array values

// Given an array: arr

var prefixarr []int

// Time Complexity: O(n), to build a prefix sum, iterating each values of the original array
// Space Complexity: O(n), to build a prefix sum, additional data structure, if we don't need
// the original array, we can overwrite it with its prefix sum, which will bring the space
// complexity down from O(n) to O(1)
func PrefixSum(arr []int) []int {
	total := 0
	for _, v := range arr {
		total += v
		prefixarr = append(prefixarr, total)
	}
	return prefixarr
}

// PostfixSum
// rangeSum

// Time Complexity: O(1)
// Space Complexity: O(1)
func rangeSum(a, b int) int {
	var preLeft int
	if a > 0 {
		preLeft = prefixarr[a-1]
	} else {
		preLeft = 0
	}
	preRight := prefixarr[b]
	return preRight - preLeft
}
