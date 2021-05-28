package slidingwindow

func minOperations(nums []int, x int) int {
	l, r := 0, 0
	res := -1
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum < x {
		return -1
	}

	cur := 0
	for r < len(nums) {
		cur += nums[r]
		for cur > sum-x {
			cur -= nums[l]
			l++
		}
		if cur == sum-x {
			res = max(res, r-l+1)
		}
		r++
	}

	if res == -1 {
		return -1
	}
	return len(nums) - res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
