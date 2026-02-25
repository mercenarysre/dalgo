// an edge case is if the length of the string is zero
// i want to start by declaring a hashmap which has a key type of byte
// and a value type of struct{}, i will iterate through each character
// in the string and add the character to a hashmap, everytime i add a
// character to the hashmap, i will increase a defined LENGTH variable
// by 1, the length variable here is the length of the substring so far
// as i iterate through the characters of the string, in a case whereby
// i come across a character which is already in the hashmap----in this manner
// we have a duplicate character, i want to break the loop cycle and return
// the length of the longest substring so far
package main

// Time Complexity: O(n)
// Space Complexity: O(k), k is the number of unique characters in the string
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	L := 0
	var length int
	window := make(map[byte]struct{})

	for R := 0; R < len(s); R++ {
		if _, ok := window[s[R]]; ok {
			/*
				for key := range m {
					delete(m, key)
				}
			*/
			clear(window)
			L = R + 1
			length = 0
			continue
		}

		window[s[R]] = struct{}{}
		length = R - L + 1
	}
	return length
}
