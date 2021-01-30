package backtrace

func permute(nums []int) (res [][]int) {
	n := len(nums)
	path := []int{}
	visited := make([]bool, n)
	var backTrace func(idx int)
	backTrace = func(idx int) {
		if idx == n {
			res = append(res, append([]int(nil), path...))
			return
		}
		for i := 0; i < n; i++ {
			if visited[i] {
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
