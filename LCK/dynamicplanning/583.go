package dynamicplanning

//计算出最长公共子序列长度dp[m][n], 最少删除次数即为 (m + n - 2*dp[m][n])
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
				continue
			}
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return (m + n - 2*dp[m][n])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
