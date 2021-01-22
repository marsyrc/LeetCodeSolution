package dynamicplanning

import "math"

//买股票，有手续费fee
func maxProfit(prices []int, fee int) int {
	dpI0 := 0
	dpI1 := math.MinInt64
	for _, v := range prices {
		dpI0, dpI1 = max(dpI0, dpI1+v), max(dpI1, dpI0-v-fee)
	}
	return dpI0
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
