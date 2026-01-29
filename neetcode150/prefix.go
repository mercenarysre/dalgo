// i initialize the prefix as the first string in the array, i want to loop over the other strings and
// i want to make a condition to compare the characters (index) of each string with the CURRENT
// prefix characters, taking into cognizance the length of the prefix and length of each string,
// for the index where the character does not match, i break and shrink the prefix based on the length of the condition
// for subsequent strings i want to compare the CURRENT prefix to the
// strings characters, if they match the prefix stays the same, if they don't i break the cycle
// after processing all strings, i return the remaining prefix

// Time Complexity: O(n * m), n is the length of the shortest string, m is the number of strings
// Space Complexity: O(1), since we did not use extra space
package main

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		j := 0
		for j < len(prefix) && j < len(strs[i]) {
			if prefix[j] != strs[i][j] {
				break
			}
			j++
		}
		prefix = prefix[:j]
	}
	return prefix
}
