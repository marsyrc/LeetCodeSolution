package dynamicplanning

func minDifficulty(jobDifficulty []int, d int) int {
	n := len(jobDifficulty)

	if n < d {
		return -1
	}

	maxV := make([][]int, n)
	for i := range maxV {
		maxV[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if i == j {
				maxV[i][j] = jobDifficulty[i]
			} else {
				maxV[i][j] = max(jobDifficulty[j], maxV[i][j-1])
			}
		}
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, d)
		for j := range dp[i] {
			dp[i][j] = 100000
		}
	}
	for i := 0; i < d; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				dp[j][i] = maxV[0][j]
			} else {
				for k := 0; k < j; k++ {
					dp[j][i] = min(dp[k][i-1]+maxV[k+1][j], dp[j][i])
				}
			}
		}
	}
	return dp[n-1][d-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
