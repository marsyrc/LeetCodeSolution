package depthfirstsearch

func findTargetSumWays(nums []int, target int) int {
	res := 0
	var dfs func(cur int, idx int)
	dfs = func(cur int, idx int) {
		if idx == len(nums) {
			if cur == target {
				res++
			}
			return
		}
		dfs(cur+nums[idx], idx+1)
		dfs(cur-nums[idx], idx+1)
	}
	dfs(0, 0)
	return res
}
