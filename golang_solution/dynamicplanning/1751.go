package dynamicplanning

import (
	"math"
	"sort"
)

func maxValue(events [][]int, k int) int {
	sort.Slice(events, func(i, j int) bool {
		return events[i][1] < events[j][1]
	})
	n := len(events)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, k+1)
		for j := range dp[i] {
			dp[i][j] = math.MinInt32
		}
	}
	dp[0][0] = 0
	dp[0][1] = events[0][2]
	for i := 1; i < n; i++ {
		l, r := 0, i
		for r-l > 1 {
			mid := (l + r) / 2
			if events[mid][1] >= events[i][0] {
				r = mid
			} else {
				l = mid
			}
		}
		if events[l][1] < events[i][0] {
			for j := 1; j <= k; j++ {
				dp[i][j] = max(dp[i][j], dp[l][j-1]+events[i][2])
			}
		} else {
			dp[i][1] = max(dp[i][1], events[i][2])
		}
		for j := 0; j <= k; j++ {
			dp[i][j] = max(dp[i][j], dp[i-1][j])
		}
	}
	res := 0
	for i := 0; i <= k; i++ {
		res = max(res, dp[n-1][i])
	}
	return res
}

func max(a ...int) int {
	res := math.MinInt32
	for _, v := range a {
		if v > res {
			res = v
		}
	}
	return res
}
