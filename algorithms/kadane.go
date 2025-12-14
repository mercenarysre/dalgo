package main

// Time Complexity: O(n>2)
func bruteForce(arr []int) int {
	maxSum := arr[0]

	for i := 0; i < len(arr); i++ {
		curSum := 0
		for j := i; j < len(arr); j++ {
			curSum += arr[j]
			if curSum > maxSum {
				maxSum = curSum
			}
		}
	}
	return maxSum
}

// Time Complexity: O(n)
// Space Complexity: O(1)
// Kadane's Algorithm
// Finding the maximum sum of the non-empty subarrays of an array
// if all elements in the array are positive, the maximum sum subarray is the entire array
// if all elements in the array are positive and negative numbers, kadane the maxium sum subarray
// if all elements in the array are negative, kadane the maximum subarray which is the highest negative number
// At each index i, what is the maximum sum subarray that ends at index i
func kadanes(arr []int) int {
	maxSum := arr[0]
	curSum := 0

	for _, v := range arr {
		// if curSum >= 0, extend the maximum sum subarray
		// if curSum < 0, reset to 0 as it carries negative baggage into next candidate subarray
		curSum = max(curSum, 0)
		curSum += v
		maxSum = max(maxSum, curSum)
	}
	return maxSum
}

func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}

// Return the left and right index of the maximal sum subarray, by keeping track of a window;
// A window denotes the contiguous subarray that does not break the constraint of the sum staying positive;
// if current sum becomes negative, we move left pointer L all the way to the right pointer R, the constraint
// is broken and we remove all elemets from the left and start a new window
func slidingWindow(arr []int) (int, int, int) {
	maxSum := arr[0]
	curSum := 0

	// maxL and maxR keep track of the subarray that contains the maximum sum
	maxL, maxR := 0, 0
	L := 0

	for R := 0; R < len(arr); R++ {
		// if current sum becomes negative, we move left pointer L all the way to the right pointer R,
		// the constraint is broken and we remove all elemets from the left and start a new window
		if curSum < 0 {
			curSum = 0
			L = R
		}

		curSum += arr[R]

		if curSum > maxSum {
			maxSum = curSum
			maxL, maxR = L, R
		}
	}
	return maxL, maxR, maxSum
}
