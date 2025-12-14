package main

// Sorting a set of values/datatypes; as long as there is a way to compare two values
// sorting algorithms generally work on any data type  [2, 3, 4, 1, 6]: iterate four times for outer loop
// i and j are pointers
// Time Complexity: O(n2)
// Space Complexity: O(1)
func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		j := i - 1
		for j >= 0 && arr[j+1] < arr[j] {
			tmp := arr[j+1]
			arr[j+1] = arr[j]
			arr[j] = tmp
			j--
		}
	}
	return arr
}

func insertionSort1(arr []int) []int {
	// i: second pointer
	for i := 1; i < len(arr); i++ {
		// j: first pointer
		j := i - 1
		for j >= 0 && arr[j+1] < arr[j] {
			// assign temporary variable and swap values
			tmp := arr[j+1]
			arr[j+1] = arr[j]
			arr[j] = tmp
			j--
		}
	}
	return arr
}

func insertionSort2(arr []int) []int {
	// i: first pointer
	for i := 1; i < len(arr); i++ {
		// j: second pointer
		j := i - 1
		for j >= 0 && arr[j+1] < arr[j] {
			// assign temporary variable and swap values
			tmp := arr[j+1]
			arr[j+1] = arr[j]
			arr[j] = tmp
			j--
		}
	}
	return arr
}
