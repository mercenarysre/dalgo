// an edge case is if the length of the string is equal to zero or one
// i want to start by iterating through the characters of the string
// using two pointers, a pointer L and a pointer R, both pointers starting
// at 0, if the string value of both pointers are equal, i want to add it
// to a hashmap, i want to maintain the left pointer and continue iterating
// if the string value of L is not equal to the string value of R, i want to
// declare a variable res such that for that variable less than or equal to
// k, i want to replace the current string value of the R pointer with the
// value of the L pointer; after doing k replacements i want to get the current
// length by subtracting the left pointer from the right pointer + 1
package main

import "strings"

func characterReplacement(s string, k int) int {
	if len(s) == 0 {
		return 0
	}

	if len(s) == 1 {
		return 1
	}

	L := 0
	res := 1
	var res1 []string
	for R := 0; R < len(s); R++ {
		if s[L] != s[R] && res <= k {
			s = strings.Replace(s, string(s[R]), string(s[L]), res)
			res++
		}

		if s[L] == s[R] {
			res1 = append(res1, string(s[R]))
		}

	}
	return len(res1)
}
