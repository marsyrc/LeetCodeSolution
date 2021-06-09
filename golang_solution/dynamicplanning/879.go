package dynamicplanning

func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
	m := len(group)
	dp := make([][][]int, m+1)
	for i := range dp {
		dp[i] = make([][]int, n+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, minProfit+1)
		}
	}
	const mod = int(1e9 + 7)
	for i := 1; i <= m; i++ {
		gNeed := group[i-1]
		pMade := profit[i-1]
		for j := 0; j <= n; j++ {
			for k := 0; k <= minProfit; k++ {
				if j < gNeed {
					dp[i][j][k] = dp[i-1][j][k]
				} else {
					dp[i][j][k] = (dp[i-1][j][k] + dp[i-1][j-gNeed][max(0, k-pMade)]) % mod
				}
			}
		}
	}
	res := 0
	for j := 0; j <= n; j++ {
		res += dp[m][j][minProfit]
		res %= mod
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
