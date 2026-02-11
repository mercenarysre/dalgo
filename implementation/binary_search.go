package main 

func binarySearch(arr []int, target int) int {
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

func binarySearch1(arr []int, target int) int {
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

func binarySearch2(arr []int, target int) int {
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

func binarySearch3(arr []int, target int) int {
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

func binarySearch5(arr []int, target int) int {
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

func binarySearch6(arr []int, target int) int {
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

func binarySearchRange1(low, high int) int {
	var mid int

	for low <= high {
		mid = (low + high)/2
		if mid > 10 {
			high = mid - 1 
		} else if mid < 10 {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func binarySearchRange2(low, high int) int {
	var mid int 

	for low <= high {
		mid = (low + high)/2
		if mid > 10 {
			high = mid - 1
		} else if mid < 10 {
			low = mid + 1
		} else {
			return mid 
		}
	}
	return -1 
}

func binarySearchRange3(low, high int) int {
	var mid int 

	for low <= high {
		mid = (low + high)/2
		if mid > 10 {
			high = mid - 1
		} else if mid < 10 {
			low = mid + 1 
		} else {
			return mid 
		}
	}
	return -1
}

func binarySearchRange4(low, high int) int {
	var mid int 

	for low <= high {
		mid = (low + high)/2 
		if mid > 10 {
			high = mid - 1
		} else if mid < 10 {
			low = mid + 1
		} else {
			return mid 
		}
	}
	return -1
}

for binarySearchRange5(low, high int) int {
	var mid int 

	for low <= high {
		mid = (low + high)/2
		if mid > 10 {
			high = mid - 1 
		} else if mid < 10 {
			low = mid + 1 
		} else {
			return mid 
		}
	}
	return -1
}