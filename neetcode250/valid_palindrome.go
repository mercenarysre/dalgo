// an edge case is if the length of the string is 1, i return false, it is not a palindrome
// i start by defining my left pointer and right pointer, start a loop such that
// the left pointer is less than the right pointer, if the array value of the left pointer
// is not equal to the array value of the right pointer i return false it is not a palindrome '
// otherwise i increment the left pointer and the right pointer, the loop starts again
// i want to ignore non-alphanumeric characters using regexp

package main

import (
	"fmt"
	"regexp"
	"strings"
)

// Time Complexity: O(n)
// Space Compleixty: O(1)
func isPalindrome(s string) bool {
	if len(s) == 0 || len(s) == 1 {
		return false
	}

	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		fmt.Println(err)
	}
	processedString := reg.ReplaceAllString(s, "")
	result := strings.ToLower(processedString)

	L := 0
	R := len(result) - 1

	for L < R {
		if result[L] != result[R] {
			return false
		}
		L++
		R--
	}
	return true
}
