package dynamicplanning

import "math"

func stoneGameIII(stoneValue []int) string {
	n := len(stoneValue)
	prefix := make([]int, n+1)
	sum := 0
	for i := 1; i <= n; i++ {
		prefix[i] = prefix[i-1] + stoneValue[i-1]
		sum += stoneValue[i-1]
	}

	dp := make([]int, n+1)
	dp[n] = 0
	for i := n - 1; i >= 0; i-- {
		d := math.MaxInt64
		for j := 1; j <= 3 && i+j <= n; j++ {
			d = min(dp[i+j], d)
		}
		dp[i] = sum - prefix[i] - d
	}
	if dp[0]*2 > sum {
		return "Alice"
	} else if dp[0]*2 == sum {
		return "Tie"
	}
	return "Bob"
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
