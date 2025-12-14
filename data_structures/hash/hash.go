package main

// A hashmap is a list of paired values, the first item in each pair is called a key
// the second item is called a value, the key and value has significant association
// with each other, not naturally paired data
// Each key can only exist once in the hash table, but there can be multiple instances of a value
// Questions across "count", "unique", "frequency"
// Time Complexity: O(n), where n is reprsenting the number of insertions when the hashmap is implemented.
// Space Complexity: O(n), where n is proportional to the number of unique keys in the array.
import (
	"fmt"
)

func main() {
	names := []string{"alice", "brad", "collin", "brad", "dylan", "kim"}
	countMap := make(map[string]int)

	for _, name := range names {
		// checking if name exists in the hashmap
		_, ok := countMap[name]
		// if name does not exists | exists
		if !ok {
			countMap[name] = 1
		} else {
			countMap[name]++
		}
	}
	fmt.Println(countMap)
}
