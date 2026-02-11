package implementation

// import "fmt"

// Prefix array
var PrefixArr []int

func generatePrefix(arr []int) []int {
	total := 0
	for _, v := range arr {
		total = total + v
		PrefixArr = append(PrefixArr, total)
	}
	return PrefixArr
}

func rangeSum(a, b int) int {
	var prefLeft int
	if a > 0 {
		prefLeft = PrefixArr[a-1]
	} else {
		prefLeft = 0
	}
	preRight := PrefixArr[b]
	return preRight - prefLeft
}

var PrefixArr1 []int

func generatePrefix1(arr []int) []int {
	total := 0
	for _, v := range arr {
		total = total + v
		PrefixArr1 = append(PrefixArr1, total)
	}
	return PrefixArr1
}

func rangeSum1(a, b int) int {
	var prefLeft int
	if a > 0 {
		prefLeft = PrefixArr1[a-1]
	} else {
		prefLeft = 0
	}
	preRight := PrefixArr1[b]
	return preRight - prefLeft
}

var PrefixArr2 []int

func generatePrefix2(arr []int) []int {
	total := 0
	for _, v := range arr {
		total = total + v
		PrefixArr2 = append(PrefixArr2, total)
	}
	return PrefixArr2
}

func rangeSum2(a, b int) int {
	var PreLeft int
	if a > 0 {
		PreLeft = PrefixArr2[a-1]
	} else {
		PreLeft = 0
	}
	preRight := PrefixArr2[b]
	return preRight - PreLeft
}

var PrefixArr3 []int

func generatePrefix3(arr []int) []int {
	total := 0
	for _, v := range arr {
		total = total + v
		PrefixArr3 = append(PrefixArr3, total)
	}
	return PrefixArr3
}

func rangeSum3(a, b int) int {
	var PreLeft int
	if a > 0 {
		PreLeft = PrefixArr3[a-1]
	} else {
		PreLeft = 0
	}
	PreRight := PrefixArr3[b]
	return PreRight - PreLeft
}

var PrefixArr4 []int

func generatePrefix4(arr []int) []int {
	total := 0
	for _, v := range arr {
		total = total + v
		PrefixArr4 = append(PrefixArr4, total)
	}
	return PrefixArr4
}

func rangeSum4(a, b int) int {
	var PreLeft int
	if a > 0 {
		PreLeft = PrefixArr4[a-1]
	} else {
		PreLeft = 0
	}
	PreRight := PrefixArr4[b]
	return PreRight - PreLeft
}

var PrefixArr5 []int

func generatePrefix5(arr []int) []int {
	total := 0
	for _, v := range arr {
		total = total + v
		PrefixArr5 = append(PrefixArr5, total)
	}
	return PrefixArr5
}

func rangeSum(a, b int) int {
	var PreLeft int
	if a > 0 {
		PreLeft = PrefixArr5[a-1]
	} else {
		PreLeft = 0
	}
	PreRight := PrefixArr5[b]
	return PreRight - PreLeft
}
