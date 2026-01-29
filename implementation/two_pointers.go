package implementation

// sorted array
func targetSum(arr []int, target int) (int, int) {
	L := 0
	R := len(arr) - 1

	for L < R {
		if arr[L]+arr[R] > target {
			R--
		} else if arr[L]+arr[R] < target {
			L++
		} else {
			return L, R
		}
	}
	return -1, -1
}

// assumes array is sorted
func targetSum1(arr []int, target int) (int, int) {
	L := 0
	R := len(arr) - 1

	for L < R {
		if arr[L]+arr[R] > target {
			R--
		} else if arr[L]+arr[R] < target {
			L++
		} else {
			return L, R
		}
	}
	return -1, -1
}

// assumes array is sorted
func targetSum2(arr []int, target int) (int, int) {
	L := 0
	R := len(arr) - 1

	for L < R {
		if arr[L]+arr[R] > target {
			R--
		} else if arr[L]+arr[R] < target {
			L++
		} else {
			return L, R
		}
	}
	return -1, -1
}

// assumes array is sorted
func targetSum3(arr []int, target int) (int, int) {
	L := 0
	R := len(arr) - 1

	for L < R {
		if arr[L]+arr[R] > target {
			R--
		} else if arr[L]+arr[R] < target {
			L++
		} else {
			return L, R
		}
	}
	return -1, -1
}
