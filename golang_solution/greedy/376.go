package greedy

//摆动序列长度
//贪心策略，最后一个下降的长度是最后一个上升的长度 + 1
func wiggleMaxLength(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	incr, decr := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			incr = decr + 1
		} else if nums[i] < nums[i-1] {
			decr = incr + 1
		}
	}

	return max(incr, decr)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
