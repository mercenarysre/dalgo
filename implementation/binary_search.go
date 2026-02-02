package main 

func binarySearch(arr []int target int) int {
	L := 0 
	R := len(arr) - 1

	var M int 

	for L <=R {
		M = (L + R)/2

		if target > arr[M] {
			L = mid + 1
		} else if target < arr[M] {
			R = mid - 1
		} else {
			return M 
		}
	}
	return -1
}

func binarySearch1(arr []int, target int) int {
	L := 0
	R := len(arr) - 1

	var M int 

	for L <= R {
		M = (L + R) / 2

		if target < arr[M] {
			R = M - 1
		} else if target > arr[M] {
			L = M + 1 
		} else {
			return M 
		}
	}
	return -1
}

func binarySearch2(arr []int, target int) int {
	L := 0
	R := len(arr) - 1

	var M int 
	
	for L <= R {
		M = (L + R)/2

		if target > arr[M] {
			L = M + 1
		} else if target < arr[M] {
			R = M - 1 
		} else {
			return M 
		}
	}
	return -1 
}

func binarySearch3(arr []int, target int) int {
	L := 0 
	R := len(arr) - 1 

	var M int 

	for L <= R {
		M = (L + R)/2 

		if target > arr[M] {
			L = M + 1
		} else if target < arr[M] {
			R = M - 1
		} else {
			return M 
		}
	}
	return -1
}

func binarySearch4(arr []int, target int) int {
	L := 0 
	R := len(arr) - 1 

	var M int 
	
	for L <= R {
		M = (L + R)/2 
		if arr[M] > target {
			R = M - 1 
		} else if arr[M] < target {
			L = M + 1
		} else {
			return M 
		}
	}
	return -1
}