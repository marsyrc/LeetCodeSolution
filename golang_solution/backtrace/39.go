package backtrace

func combinationSum(candidates []int, target int) [][]int {
	path := []int{}
	res := [][]int{}
	var dfs func(idx int, target int)
	dfs = func(idx int, target int) {
		if idx == len(candidates) {
			return
		}
		if target == 0 {
			res = append(res, append([]int(nil), path...))
			return
		}
		dfs(idx+1, target)

		if target-candidates[idx] >= 0 {
			path = append(path, candidates[idx])
			dfs(idx, target-candidates[idx])
			path = path[:len(path)-1]
		}
	}
	dfs(0, target)
	return res
}
