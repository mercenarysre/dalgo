// edge case i want to check is if the string has length 0 or length 1,
// then it is not a palindrome, return false
// i want to check the string if it is a palindrome, i do this by defining left
// and right pointers, left pointer start at the first element of the string,
// right pointer start at the last element of the string, then i define a loop
// such that the left pointer is less than the right pointer i check if the string
// value of the left pointer is not equal to the string value of the right pointer
// it is not a palindrome, given the string is not a palindrome now for the consideration of deleting at most on character
// from the string, the intuition here is that the right pointer value is what does not
// make it a palindrome, because it's just common sense, the right pointer string value has to be
// the same as the left pointer string value, given where i am for the right pointer string
// value in the loop above when it is not palindrome, i want to remove the current right pointer
// string value and replace it with the next string value after it, the length of the string
// is reduced and i want to decrease the value of the left pointer and starts the loop again
// if its not a palindrome break and return false, otherwise return true

package main

func validPalindrome(s string) bool {
	if len(s) == 0 || len(s) == 1 {
		return false
	}

	L := 0
	R := len(s) - 1

	for L < R {
		if s[L] != s[R] {
			if L > 0 {
				L--
			} else {
				L = 0
			}
			if R == len(s)-1 {
				s = s[:R]
				R = len(s) - 1
			} else {
				s = s[:R] + s[R+1:]
			}
			if s[L] != s[R] {
				return false
			}
		}
		L++
		R--
	}
	return true
}
