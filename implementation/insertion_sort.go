package main

func insertionSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

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
	if len(arr) <= 1 {
		return arr
	}

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

func insertionSort2(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

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

func insertionSort3(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

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

func insertionSort4(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

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

func insertion5(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

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

func insertionSort6(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

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

func insertionSort7(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

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

func insertionSort8(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	for i := 0; i < len(arr); i++ {
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

func insertionSort9(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

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
