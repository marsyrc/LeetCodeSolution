package dynamicplanning

import "math"

func lenLongestFibSubseq(arr []int) int {
	//special case
	if len(arr) == 0 {
		return 0
	}

	//init dp
	dp := make([][]int, len(arr))
	for i := range dp {
		dp[i] = make([]int, len(arr))
	}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			dp[i][j] = 2
		}
	}

	//store the pairs of(arr[i], i)
	index := make(map[int]int)
	for i, v := range arr {
		index[v] = i
	}
	res := 0
	for k := 0; k < len(arr); k++ {
		for j := 0; j < k; j++ {
			//if i < j && (arr[j] - arr[k]) isExist in arr
			if i, ok := index[arr[k]-arr[j]]; i < j && ok {
				dp[j][k] = dp[i][j] + 1
			}
			res = max(res, dp[j][k])
		}
	}

	//illegal
	if res < 3 {
		return 0
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
