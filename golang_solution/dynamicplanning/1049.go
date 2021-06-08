package dynamicplanning

func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, stone := range stones {
		sum += stone
	}
	target := sum / 2
	dp := make([]int, target+1)
	for _, stone := range stones {
		for i := target; i >= stone; i-- {
			dp[i] = max(dp[i], dp[i-stone]+stone)
		}
	}
	return sum - 2*dp[target]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
