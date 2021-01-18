package backtrace

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	path := []int{}
	sort.Ints(candidates)
	var dfs func(idx int, target int)
	dfs = func(idx int, target int) {
		if target == 0 {
			res = append(res, append([]int(nil), path...))
			return
		}
		for i := idx; i < len(candidates); i++ {
			if target-candidates[i] < 0 {
				break
			}
			//剪枝：去掉同层不是第一个的相同值
			if i > idx && candidates[i-1] == candidates[i] {
				continue
			}
			path = append(path, candidates[i])
			dfs(i+1, target-candidates[i])
			path = path[:len(path)-1]
		}
	}
	dfs(0, target)
	return res
}
