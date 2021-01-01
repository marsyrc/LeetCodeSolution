package dynamicplanning

import "math"

//convert this problem to rob(198.go, 213.go)
func deleteAndEarn(nums []int) int {
	//record the count of index(of nums)
	cnt := make([]int, 10001)

	//special case
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	//record the max Value in nums
	maxN := -1
	for _, v := range nums {
		cnt[v]++
		maxN = max(maxN, v)
	}

	dp := make([]int, maxN+1)
	dp[1] = cnt[1]
	dp[2] = max(dp[1], cnt[2]*2)
	for i := 2; i <= maxN; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+cnt[i]*i)
	}
	return dp[maxN]
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
