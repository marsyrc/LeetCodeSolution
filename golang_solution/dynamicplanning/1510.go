package dynamicplanning

func winnerSquareGame(n int) bool {
	dp := make([]bool, n+1)
	for i := 1; i <= n; i++ {
		for k := 1; k*k <= i; k++ {
			if !dp[i-k*k] {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}
