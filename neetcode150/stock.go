// i want to start by defining a maxprofit variable, then a left pointer which starts at 0
// then i want to define a for loop which iterates over the integer array, the for loop
// uses a right pointer which starts at 1 and is less than the length of the integer
// array, in the for loop i want to declare a profit variable which is the subtraction
// between the array value of the right pointer and the left pointer
// if the profit variable is less than or equal to zero, the transactipm is
// not profitable, i want to shrink the window and increase the left pointer
// if the profit variable is less than maxprofit, then maxprofit is still
// equal to maxprofit, else if the profit varible is greater then or equal to
// maxprofit, maxprofit is equal to profit variable
// return maxprofit or return zero

package main

func maxProfit(prices []int) int {
	var maxprofit int
	var profit int
	L := 0

	if len(prices) <= 1 {
		return 0
	}

	for R := 1; R < len(prices); R++ {
		profit = prices[R] - prices[L]

		if profit <= 0 {
			L++
		} else if profit < maxprofit {
			maxprofit = maxprofit
		} else if profit >= maxprofit {
			maxprofit = profit
		}
	}
	if maxprofit > 0 {
		return maxprofit
	}
	return 0
}
