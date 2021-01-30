package backtrace

import "sort"

func permuteUnique(nums []int) (res [][]int) {
	n := len(nums)
	path := []int{}
	visited := make([]bool, n)
	sort.Ints(nums)
	var backTrace func(idx int)
	backTrace = func(idx int) {
		if idx == n {
			res = append(res, append([]int(nil), path...))
			return
		}
		for i := 0; i < n; i++ {
			//如果前面有重复元素使用过并回溯了，就剪枝避免重复解
			if visited[i] || (i > 0 && !visited[i-1] && nums[i] == nums[i-1]) {
				continue
			}
			visited[i] = true
			path = append(path, nums[i])
			backTrace(idx + 1)
			path = path[:len(path)-1]
			visited[i] = false
		}
	}
	backTrace(0)
	return res
}
