// an edge case is if the matrix array is empty, return false
// i am going to write a brute force approach for the matrix problem
// such that i want to start a for loop which starts from the matrix array
// itself, inside that for loop i want to start
// another for loop for the first (and subsequent) elements (subarrays),
// i want to iterate the elements of the subarray, if any of the elements
// is equal to the target--return true, if not return false

// Time Complexity: O(n)
// Space Complexity: O(1)
package main

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == target {
				return true
			}
		}
	}
	return false
}
