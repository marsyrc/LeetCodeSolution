package backtrace

//子集,无重复情况
func subsets(nums []int) [][]int {
	res := [][]int{}
	if len(nums) == 0 {
		return res
	}
	n := len(nums)
	path := []int{}
	var backTrace func(idx int)
	backTrace = func(idx int) {
		res = append(res, append([]int(nil), path...))
		for i := idx; i < n; i++ {
			path = append(path, nums[i])
			backTrace(i + 1)
			path = path[:len(path)-1]
		}
	}
	backTrace(0)
	return res
}
