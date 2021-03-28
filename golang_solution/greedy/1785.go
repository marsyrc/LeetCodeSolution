package greedy

func minElements(nums []int, limit int, goal int) int {
	sum := goal
	for i := range nums {
		sum -= nums[i]
	}
	if sum == 0 {
		return 0
	}
	sum = abs(sum)
	remain := sum % abs(limit)
	res := sum / abs(limit)
	if remain != 0 {
		return res + 1
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
