package dynamicplanning

func numTilings(N int) int {
	dp := make([]int, N+1)
	mod := 1000000000 + 7
	dp[0] = 1
	dp[1] = 1
	dp[2] = 2
	dp[3] = 5

	//f[n] = f[n - 1] + f[n - 2] + 2 * (f[n - 3] + f[n - 4] + ... + f[0])
	//which equals to f[n] = 2 * f[n - 1] + f[n -3]

	for i := 4; i <= N; i++ {
		dp[i] = (2*(dp[i-1]%mod))%mod + (dp[i-3] % mod)
	}
	return dp[N] % mod
}
