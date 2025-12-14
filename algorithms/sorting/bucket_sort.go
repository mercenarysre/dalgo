package main

// Algorithm to use if all of the range of values in the provided
// Assuming all values in the input are in a specified range [ 2, 2, 1, 0, 1, 0]
// Time Complexity: O(n), case of nested loop where the outer loop iterates n times, where n is the length of the array
// however the inner loop runs for counts[n] which is different everytime and cannot be factored into the time complexity,
// thus, O(n)
// Space Complexity: O(1)
func bucketSort(arr []int) []int {
	// Assuming arr only contains 0, 1 or 2
	counts := [3]int{0, 0, 0}

	// Count the quantity of each val in arr
	for _, num := range arr {
		counts[num]++
	}

	// Fill each bucket in the original array
	i := 0
	for n, count := range counts {
		for j := 0; j < count; j++ {
			arr[i] = n
			i++
		}
	}
	return arr
}

func bucketSort1(arr []int) []int {
	counts := [4]int{0, 0, 0, 0}

	for _, num := range arr {
		counts[num]++
	}

	i := 0
	for n, count := range counts {
		for j := 0; j < count; j++ {
			arr[i] = n
			i++
		}
	}
	return arr
}

func bucketSort2(arr []int) []int {
	counts := [5]int{0, 0, 0, 0, 0}

	for _, num := range counts {
		counts[num]++
	}

	i := 0
	for n, count := range counts {
		for j := 0; j < count; j++ {
			arr[i] = n
			i++
		}
	}
	return arr
}
