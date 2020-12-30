package dynamicplanning

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	//ergodic every j in [1, i], use it to be root
	//[1, j - 1] as leftNode, [j + 1, n] as rightNode
	//dp[i] is a Cartesian product
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
