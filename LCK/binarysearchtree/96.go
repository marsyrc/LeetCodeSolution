package binarysearchtree

//dp solution:
//let j in [1, n] be the current root,
//so dp[1, i] = dp[1, j - 1] + dp[j + 1, i] = dp[1, j - 1] + dp[1, i - j].
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
