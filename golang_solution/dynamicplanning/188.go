package dynamicplanning

import "math"

func maxProfit(k int, prices []int) int {
	//k too big = inf chances of purchasing
	if k > len(prices)/2 {
		return maxProfitInf(prices)
	}
	//dp_i_1 : current holding tickets
	//dp_i_0 : current not holding tickets
	dp_i_1, dp_i_0 := make([]int, k+1), make([]int, k+1)
	for i := 0; i <= k; i++ {
		dp_i_0[i] = 0
		dp_i_1[i] = math.MinInt16
	}

	for _, v := range prices {
		for i := k; i >= 1; i-- {
			dp_i_0[i] = max(dp_i_0[i], dp_i_1[i]+v)
			dp_i_1[i] = max(dp_i_1[i], dp_i_0[i-1]-v)
		}
	}
	return dp_i_0[k]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfitInf(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	sum := 0
	n := len(prices)
	for i := 1; i < n; i++ {
		if prices[i]-prices[i-1] > 0 {
			sum += prices[i] - prices[i-1]
		}
	}
	return sum
}
