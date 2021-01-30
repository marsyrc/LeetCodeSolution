package dynamicplanning

import "math"

//0-1 package problem
func canPartition(nums []int) bool {
	max := math.MinInt32
	sum := 0
	for _, v := range nums {
		sum += v
		if v > max {
			max = v
		}
	}

	target := sum / 2
	if max > target {
		return false
	}

	if sum%2 != 0 {
		return false
	}

	//general dp[i][j] : can find some Ints (sum is j) from [0, i]
	//so dp[i][j] = dp[i - 1][j] || dp[i - 1][j - nums[i]]  (choose nums[i] or not)
	dp := make([][]bool, len(nums))
	for i := range dp {
		dp[i] = make([]bool, target+1)
	}

	if nums[0] <= target {
		dp[0][nums[0]] = true
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j <= target; j++ {
			dp[i][j] = dp[i-1][j]
			if nums[i] == j {
				dp[i][j] = true
			}
			if nums[i] < j {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
			}
		}
	}
	return dp[len(nums)-1][target]

	//dp has been condense
	// dp := make([]bool, target+1)
	// dp[0] = true
	// for i := 0; i < len(nums); i++ {
	// 	v := nums[i]
	// 	for j := target; j >= v; j++ {
	// 		if dp[target] == true {
	// 			return true
	// 		}
	// 		dp[j] = dp[j] || dp[j-v]
	// 	}
	// }
	// return dp[target]
}
