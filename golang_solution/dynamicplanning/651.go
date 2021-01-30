package dynamicplanning

func maxA(N int) int {
	dp := make([]int, N+1)
	for i := 1; i <= N; i++ {
		dp[i] = dp[i-1] + 1
		for j := 0; j < i-1; j++ {
			dp[i] = max(dp[i], dp[j]*(i-j-1))
		}
	}
	return dp[N]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
