// an edge case i want to consider is if the length of the array is 0 or 1
// i want to start by defining left and right pointer for the array, such that
// the left pointer start at the first element of the array, and the right pointer
// starts at last element of the array, then i start a loop such that the left pointer
// is less than or equal to the right pointer, the array value of the right pointer
// is put in place for the array value of the left pointer, and vice versa
// when the loop is completed, i return the array

package main

// Time Complexity: O(n)
// Space Complexity: 0(1)
func reverseString(s []byte) []byte {
	if len(s) == 0 || len(s) == 1 {
		return s
	}

	L := 0
	R := len(s) - 1

	for L <= R {
		n := s[L]
		s[L] = s[R]
		s[R] = n
		L++
		R--
	}
	return s
}
