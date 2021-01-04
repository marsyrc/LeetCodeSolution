package dynamicplanning

//双串问题，类似97.go
func numDistinct(s string, t string) int {
	m, n := len(t), len(s)
	//dp : T first i chars be made from S first j chars
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		dp[i][0] = 0
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if t[i-1] == s[j-1] {
				//two situations:
				//1. use t[i - 1] matches s[j - 1]
				//2. not use s[j - 1] but former s's element to match t[i - 1], so dp[i][j - 1]
				dp[i][j] = dp[i-1][j-1] + dp[i][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[m][n]
}
