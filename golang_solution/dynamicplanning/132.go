package dynamicplanning

import "math"

func minCut(s string) int {
	if len(s) < 2 {
		return 0
	}
	check := make([][]bool, len(s))
	for i := range check {
		check[i] = make([]bool, len(s))
	}
	//区间dp
	for l := 1; l <= len(s); l++ {
		for left := 0; left <= len(s)-l; left++ {
			right := left + l - 1
			if s[left] == s[right] && (right-left <= 2 || check[left+1][right-1]) {
				check[left][right] = true
			}
		}
	}

	//dp : [0, i] 所需的最少分割次数
	dp := make([]int, len(s))
	for i := range dp {
		dp[i] = i
	}
	for i := 1; i < len(s); i++ {
		if check[0][i] {
			dp[i] = 0
			continue
		}
		for j := 0; j < i; j++ {
			if check[j+1][i] {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}
	return dp[len(s)-1]

}

func min(a ...int) int {
	res := math.MaxInt32
	for _, v := range a {
		if res > v {
			res = v
		}
	}
	return res
}
