package main

// Time Complexity: O(n>2)
// Space Complexity: O(1)
func longest(s string) int {
	length := 0
	n := len(s)

	for i := 0; i < n; i++ {
		// odd length palindrome (center at i)
		l, r := i, i
		for l >= 0 && r < n && s[l] == s[r] {
			if r-l+1 > length {
				length = r - l + 1
			}
			l--
			r++
		}

		// even length palindrome (center between i and i+1)
		l, r = i, i+1
		for l >= 0 && r < n && s[l] == s[r] {
			if r-l+1 > length {
				length = r - l + 1
			}
			l--
			r++
		}
	}

	return length
}
