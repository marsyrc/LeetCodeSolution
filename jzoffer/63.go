package jianzhioffer

import "math"

func maxProfit(prices []int) int {
	cost := math.MaxInt32
	profit := 0
	for _, v := range prices {
		cost = min(cost, v)
		profit = max(profit, v-cost)
	}
	return profit
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
