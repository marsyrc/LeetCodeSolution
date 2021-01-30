package dynamicplanning

func maxProfit(prices []int) int {
	// 3 states
	//1. holding
	//2. no holding, in cd
	//3. no holding, not in cd
	if len(prices) == 0 {
		return 0
	}
	dp1 := -prices[0]
	dp2 := 0
	dp3 := 0

	for i := 1; i < len(prices); i++ {
		dp1, dp2, dp3 = max(dp1, dp3-prices[i]), dp1+prices[i], max(dp3, dp2)
	}
	return max(dp2, dp3)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
