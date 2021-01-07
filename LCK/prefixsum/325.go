package prefixsum

func maxSubArrayLen(nums []int, k int) int {
	sum := make([]int, len(nums)+1)
	m := make(map[int]int)
	m[0] = 0
	for i := 1; i <= len(nums); i++ {
		sum[i] = nums[i-1] + sum[i-1]
		if _, ok := m[sum[i]]; !ok {
			m[sum[i]] = i
		}
	}
	ans := 0
	for i := len(nums); i >= 0; i-- {
		if j, ok := m[sum[i]-k]; ok {
			ans = max(ans, i-j)
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
