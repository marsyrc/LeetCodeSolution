package dynamicplanning

import "math"

func maxCoins(nums []int) int {
	//add 1 in two ends of nums
	n := len(nums)
	temp := []int{}
	temp = append(temp, 1)
	temp = append(temp, nums...)
	temp = append(temp, 1)

	//dp : max res in (i, jss)
	// j = i + l - 1
	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
	}

	//ergodic every l = j - i + 1
	for l := 3; l <= n+2; l++ {
		for i := 0; i <= n+2-l; i++ {
			res := 0
			for k := i + 1; k < i+l-1; k++ {
				left := dp[i][k]
				right := dp[k][i+l-1]
				res = max(res, left+right+temp[i]*temp[k]*temp[i+l-1])
			}

			dp[i][i+l-1] = res
		}
	}
	return dp[0][n+1]
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
