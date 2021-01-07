package dynamicplanning

import "math"

//区间dp
func strangePrinter(s string) int {
	if len(s) == 0 {
		return 0
	}
	dp := make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, len(s)+1)
	}
	//boundary
	for i := 0; i <= len(s); i++ {
		dp[i][i] = 1
	}
	//l -> i
	for l := 2; l <= len(s); l++ {
		for i := 0; i < len(s); i++ {
			j := i + l - 1
			if j >= len(s) {
				break
			}
			// i 要自己涂一次，则 dp[i][j] = 1 + dp[i + 1][j]
			// 其中第一项 1 表示 i 位置单独花费一次次数
			dp[i][j] = 1 + dp[i+1][j]
			//i 与中间的某个切分位置 k 一起打印 (条件是 s[i] = s[k])，则 dp[i][j] = dp[i+1][k] + dp[k+1][j]
			//其中第一项 dp[i+1][k] 表示 i 位置跟着 k 一起转移了，不在单独考虑 i 花费的次数了
			for k := i + 1; k <= j; k++ {
				if s[k] == s[i] {
					dp[i][j] = min(dp[i][j], dp[i+1][k]+dp[k+1][j])
				}
			}
		}
	}
	return dp[0][len(s)-1]
}

func min(a ...int) int {
	res := math.MaxInt32
	for _, v := range a {
		if v < res {
			res = v
		}
	}
	return res
}
