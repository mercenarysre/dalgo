package main

// Two-branch recursive algorithm which spilts an array into sub-arrays
// recursively sort the subarrays by merging two subarrays at a time.
// two branch recursion, we solve both the branches and 'piece'
// back together the solutions to the subproblems to arrive at the ultimate solution
// The merge() call will not be executed until both the recursive mergeSort()
// calls have returned for the current subarray.
// Time Complexity: O(nlogn): having to iterate through each element in the subarrays across logn levels(n/2, height of recursive calls)
// Space Complexity: O(n): having to create temporary subarrays across logn levels
func mergeSort(arr []int, l, r int) []int {
	if l < r {
		// calculate middle index of the array
		m := (l + r) / 2

		// mergeSort split the array range into subarrays
		mergeSort(arr, l, m)
		mergeSort(arr, m+1, r)

		// merge uses the array values and puts them in order
		merge(arr, l, m, r)
	}
	return arr
}

func mergeSort1(arr []int, l, r int) []int {
	if l < r {
		m := (l + r) / 2

		// mergeSort spilt the array range into subarrays
		mergeSort1(arr, l, m)
		mergeSort1(arr, m+1, r)

		// merge uses the array values and puts them in order
		merge(arr, l, m, r)
	}
	return arr
}

func mergeSort2(arr []int, l, r int) []int {
	if l < r {
		m := (l + r) / 2

		// mergeSort spilt the array into subarrays
		mergeSort2(arr, l, m)
		mergeSort2(arr, m+1, r)

		// merge uses the array values and puts them in order
		merge(arr, l, m, r)
	}
	return arr
}

// explain the merge function
func merge(arr []int, l, m, r int) {
	length1 := m - l + 1
	length2 := r - m

	L := make([]int, length1)
	R := make([]int, length2)

	for i := 0; i < length1; i++ {
		L[i] = arr[l+i]
	}

	for j := 0; j < length2; j++ {
		R[j] = arr[m+1+j]
	}

	i := 0
	j := 0
	k := l

	for i < length1 && j < length2 {
		if L[i] <= R[j] {
			arr[k] = L[i]
			i++
		} else {
			arr[k] = R[j]
			j++
		}
		k++
	}

	for i < length1 {
		arr[k] = L[i]
		i++
		k++
	}

	for j < length2 {
		arr[k] = R[j]
		j++
		k++
	}
}
