// Programming a problem that could be solved recursively(matrix recursion)
// and ensuring that it doesn't make recursive calls for overlapping subproblems
// consider a problem where the result of subproblems depend on two variables
package main

type CountPaths struct{}

// BruteForce Time: O(2 ^ (n + m)), Space: O(n + m)
func (C *CountPaths) BruteForce(r, c, rows, cols int) int {
	if r == rows || c == cols {
		return 0
	}
	if r == rows-1 && c == cols-1 {
		return 1
	}
	return (C.BruteForce(r+1, c, rows, cols) + C.BruteForce(r, c+1, rows, cols))
}

// Memoization - Time and Space: O(n * m)
func (C *CountPaths) Memoization(r, c, rows, cols int, cache [][]int) int {
	if r == rows || c == cols {
		return 0
	}
	if cache[r][c] > 0 {
		return cache[r][c]
	}
	if r == rows-1 && c == cols-1 {
		return 1
	}
	cache[r][c] = (C.Memoization(r+1, c, rows, cols, cache) + C.Memoization(r, c+1, rows, cols, cache))
	return cache[r][c]
}

// Bottom Up - Time: O(n * m), Space: O(m), where m is num of cols
func (c *CountPaths) Dp(rows, cols int) int {
	prevRow := make([]int, cols)

	for i := rows - 1; i >= 0; i-- {
		curRow := make([]int, cols)
		curRow[cols-1] = 1
		for j := cols - 2; j >= 0; j-- {
			curRow[j] = curRow[j+1] + prevRow[j]
		}
		prevRow = curRow
	}
	return prevRow[0]
}
