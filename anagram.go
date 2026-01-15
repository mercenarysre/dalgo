// i want to check the length of both strings,
// if not the same, return false, they are not anagrams of each other
// i want to a hashmap for the first string, then iterate through the
// characters of the first string and store them(each character) in the hashmap with
// a number(occurence) in the string, i will create a second hashmap and
// store the characters of the second string in the hashmap, each character with
// a respective number(occurrence) in the string
// then i want to compare both hashmaps such that when the each first string hashmap
// character (value) does not match the second string character hashmap value, return false
// otherwise return true

// Time Complexity: O(i + j), n is the length of string s, j is the length of string j
// Space Complexity: O(1), since we have at most 26 different characters, contstant extra space
package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	n := make(map[byte]int)
	m := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		_, ok := n[s[i]]
		if ok {
			n[s[i]]++
		}
		if !ok {
			n[s[i]] = 1
		}
	}

	for i := 0; i < len(t); i++ {
		_, ok := m[t[i]]
		if ok {
			m[t[i]]++
		}
		if !ok {
			m[t[i]] = 1
		}
	}

	for k, v := range n {
		if m[k] != v {
			return false
		}
	}
	return true
}
