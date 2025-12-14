package main

import (
	"fmt"
)

// Worst Case where i = 0
func insertBeginnning(arr []int, i, n, length int) []int {
	// Make space by appending dummy element
	arr = append(arr, 0)

	// Shift starting from the end to i
	for index := length - 1; index > i-1; index-- {
		arr[index+1] = arr[index]
	}

	// Insert at i
	arr[i] = n

	// Increment length
	//	length++
	return arr
}

// Worst Case where i = 0
func removeBeginning(arr []int, i, length int) []int {
	// Shift starting from i + 1 to end
	for index := i + 1; index < length; index++ {
		arr[index-1] = arr[index]
	}
	// Shrink the array
	// arr[length-1] = 0
	// Decrement length
	// length--
	return arr[:length-1]
}

// Print array: read/iterate through values of array
func printArr(arr []int, length int) {
	for i := 0; i < length; i++ {
		// fmt.Print(arr[i], " ")
		fmt.Println(arr[i])
	}
	// fmt.Println()
}

// Searching an array for a particular value
func search(arr []int, n, length int) {
	for i := 0; i < length; i++ {
		if arr[i] == n {
			fmt.Println(i)
		} else {
			fmt.Println("The value does not exist in the array")
		}
	}
}

// Average Case
func insertMiddle(arr []int, i, n, length int) []int {
	// Shift starting from the end to i.
	for index := length - 1; index > i-1; index-- {
		arr[index+1] = arr[index]
	}
	// Insert at i
	arr[i] = n
	// Increment length
	length++
	return arr
}

// Best Case
func insertEnd(arr []int, n, length, capacity int) []int {
	if length < capacity {
		// Insert n
		arr[length] = n
		// Increment length
		length++
	}
	return arr
}

// Average Case
func removeMiddle(arr []int, i, length int) []int {
	// Shift starting from i + 1 to end.
	for index := i + 1; index < length; index++ {
		arr[index-1] = arr[index]
	}
	// No need to remove arr[i], since we already shifted
	arr[length-1] = 0
	// Decrement length
	length--
	return arr
}

// Best Case
func removeEnd(arr []int, length int) []int {
	if length > 0 {
		// Remove last element
		arr[length-1] = 0
		// Decrement length
		length--
	}
	return arr
}
