package dynamicplanning

import "math"

func maxSizeSlices(slices []int) int {
	if len(slices) == 0 {
		return 0
	}
	if len(slices) == 1 {
		return slices[0]
	}
	if len(slices) == 2 {
		return max(slices...)
	}

	length := len(slices)
	n := length / 3
	return max(byRangEat(1, length-1, slices, n), byRangEat(0, length-2, slices, n))
}

func byRangEat(start int, end int, slices []int, n int) int {
	if start == end {
		return slices[start]
	}
	k := end - start + 1
	dp := make([][]int, k)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	dp[0][0] = 0
	dp[1][0] = 0
	dp[0][1] = slices[start]
	dp[1][1] = max(slices[start+1], slices[start])
	for i := 2; i < k; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = max(dp[i-2][j-1]+slices[start+i], dp[i-1][j])
		}
	}
	return dp[k-1][n]
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
