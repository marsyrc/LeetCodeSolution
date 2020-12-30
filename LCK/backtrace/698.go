package backtrace

import "math"

func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	max := math.MinInt16
	for _, v := range nums {
		sum += v
		if v > max {
			max = v
		}
	}
	if sum%k != 0 {
		return false
	}
	sum /= k
	if max > sum {
		return false
	}
	visited := make([]bool, len(nums))

	var dfs func(k int, cur int, pos int) bool
	dfs = func(k int, cur int, pos int) bool {
		if k == 0 {
			return true
		}
		if cur == sum {
			return dfs(k-1, 0, 0)
		}

		for i := pos; i < len(nums); i++ {
			if visited[i] {
				continue
			}

			visited[i] = true
			if cur+nums[i] <= sum && dfs(k, cur+nums[i], i+1) {
				return true
			}
			visited[i] = false

		}
		return false
	}
	return dfs(k, 0, 0)
}
