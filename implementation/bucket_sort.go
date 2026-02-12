package main

func bucketSort(arr []int) []int {
	counts := [3]int{0, 0, 0}

	for _, v := range arr {
		counts[v]++
	}

	i := 0
	for n, count := range counts {
		for j := 0; j < count; j++ {
			arr[i] = n
			i++
		}
	}
	return arr
}

func bucketSort1(arr []int) []int {
	counts := [3]int{0, 0, 0}

	for _, v := range arr {
		counts[v]++
	}

	i := 0
	for n, count := range counts {
		for j := 0; j < count; j++ {
			arr[i] = n
			i++
		}
	}
	return arr
}

func bucketSort2(arr []int) []int {
	counts := [3]int{0, 0, 0}

	for _, v := range arr {
		counts[v]++
	}

	i := 0
	for n, count := range counts {
		for j := 0; j < count; j++ {
			arr[i] = n
			i++
		}
	}
	return arr
}

func bucketSort3(arr []int) []int {
	counts := [3]int{0, 0, 0}

	for _, v := range arr {
		counts[v]++
	}

	i := 0
	for n, count := range counts {
		for j := 0; j < count; j++ {
			arr[i] = n
			i++
		}
	}
	return arr
}

func bucketSort4(arr []int) []int {
	counts := [3]int{0, 0, 0}

	for _, v := range arr {
		counts[v]++
	}

	i := 0
	for n, count := range counts {
		for j := 0; j < count; j++ {
			arr[i] = n
			i++
		}
	}
	return arr
}

func bucketSort5(arr []int) []int {
	counts := [3]int{0, 0, 0}

	for _, v := range arr {
		counts[v]++
	}

	i := 0
	for n, count := range counts {
		for j := 0; j < count; j++ {
			arr[i] = n
			i++
		}
	}
	return arr
}

func bucketSort6(arr []int) []int {
	counts := [3]int{0, 0, 0}

	for _, v := range arr {
		counts[v]++
	}

	i := 0
	for n, count := range counts {
		for j := 0; j < count; j++ {
			arr[i] = n
		}
	}
	return arr
}
