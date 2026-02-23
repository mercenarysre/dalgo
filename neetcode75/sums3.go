// an edge case is if the length of the integer array is less than or equal to 1
// non-decreasing implies the array is sorted
// 1-indexed implies the integer array indices will start from 1

// using a brute force solution i want to do a loop and
// take the first element in the array and a second loop
// within it of which i will take other elements of the array
// and add it to the first element and see if it sums up to
// the target, if it does, i want to break the loop and
// return the 1-indexed indices of both elements; I will do the
// iteration for the second element, third element etc of the integer array

// i can use a binary search approach here such that i want to declare a
// left pointer and right pointer, and a mid variable, i start a loop
// such that the left pointer is always less than the right pointer
// and the mid variable is the mid(half) of the addition of the left and right
// pointers, i check if the addition of the left and pointer is greater than the
// target if it is, i want to reduce the right pointer to the length of the mid
// variable, if the addition is less than the target--i want to increasse the left
// pointer to the length of mid variable; if the addition is equal to the target
// i want to return the indices(1-indexed) of the left and right pointers
// this approach works because the array is sorted, and using common sense if the
// addition is less than the target this means the left pointer array value is just
// too small enough to meet the target and we have to increase it, same as if the right pointer
// array value is greater than the target this means the right pointer array value is
// just too large enough to meet the target and we have to reduce it, this solution
// also has a better runtime compared to the first approach

// i got the second analogy wrong, the binary search approach is for searching for a
// single value, when dealing with a pair sum, use the two pointers technique

// Time Complexity: O(n)
// Space Complexity: O(1)
package main

func twoSum(numbers []int, target int) []int {
	L := 0
	R := len(numbers) - 1
	// var M int

	for L < R {
		// M = (L + R) / 2
		if numbers[L]+numbers[R] > target {
			R--
		} else if numbers[L]+numbers[R] < target {
			L++
		} else {
			return []int{L + 1, R + 1}
		}
	}
	return []int{-1, -1}
}
