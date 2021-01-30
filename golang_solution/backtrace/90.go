package backtrace

import "sort"

func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	path := []int{}
	n := len(nums)
	sort.Ints(nums)
	var backTrace func(idx int)
	backTrace = func(idx int) {
		res = append(res, append([]int(nil), path...))
		for i := idx; i < n; i++ {
			if i > idx && nums[i] == nums[i-1] {
				continue
			}
			path = append(path, nums[i])
			backTrace(i + 1)
			path = path[:len(path)-1]
		}
	}
	backTrace(0)
	return res
}
