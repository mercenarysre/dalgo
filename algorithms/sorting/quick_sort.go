// Recursive Algorithm
package main

// Time Complexity: O(n2), here an array is not spilt into equal halves,
// there is n, n-1, n-2 levels not logn(n/2) levels
// having to iterate across each element across n levels.
// Space Complexity: O(n): having to create temporary subarrays across logn levels)
// s: pointer index
// e: pivot index
func QuickSort(arr []int, s int, e int) []int {
	// length of subarray is less than equal to 1, return array
	if e-s+1 <= 1 {
		return arr
	}

	pivot := arr[e]

	// pointer to tell which index we should insert the next value that is less than or equal to the pivot
	left := s

	for i := s; i < e; i++ {
		if arr[i] < pivot {
			tmp := arr[left]
			arr[left] = arr[i]
			arr[i] = tmp
			left++
		}
	}

	arr[e] = arr[left]
	arr[left] = pivot

	// recursive step on the left half of the array(subarray) without including the pivot
	QuickSort(arr, s, left-1)
	// recursive step on the right half of the array(subarray) without including the pivot
	QuickSort(arr, left+1, e)

	return arr
}

func QuickSort1(arr []int, s int, e int) []int {
	// length of subarray is less than equal to 1, return array
	if e-s+1 <= 1 {
		return arr
	}

	pivot := arr[e]

	// pointer to tell which index we insert the next value that is less than or equal to the pivot
	left := s

	for i := s; i < e; i++ {
		if arr[i] < pivot {
			tmp := arr[left]
			arr[left] = arr[i]
			arr[i] = tmp
			left++
		}
	}

	// making the pivot the center value
	arr[e] = arr[left]
	arr[left] = pivot

	// recursive step on the left half of the array(subarray)
	QuickSort1(arr, s, left-1)
	// recursive step on the left half of the array(subarray)
	QuickSort1(arr, left+1, e)

	return arr
}

func QuickSort2(arr []int, s int, e int) []int {
	// length of subarray is equal to 1, return array
	if e-s+1 <= 1 {
		return arr
	}

	pivot := arr[e]

	// left pointer to tell which index to insert the value that is less than or equal to the pivot
	left := s

	for i := s; i < e; i++ {
		if arr[i] < pivot {
			tmp := arr[left]
			arr[left] = arr[i]
			arr[i] = tmp
			left++
		}
	}

	arr[e] = arr[left]
	arr[left] = pivot

	// recursive step on the left half of the array(subarray) not including the pivot
	QuickSort2(arr, s, left-1)
	// recursive step on the left half of the array(subarray) not including the piv
	QuickSort2(arr, left+1, e)

	return arr
}
