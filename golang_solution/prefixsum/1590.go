package prefixsum

func minSubarray(nums []int, p int) int {
	sum := 0
	n := len(nums)
	for _, num := range nums {
		sum = (sum + num) % p
	}
	if sum == 0 {
		return 0
	}
	pos := make(map[int]int)
	cur := 0
	res := n
	pos[0] = -1
	for i := 0; i < n; i++ {
		cur += nums[i]
		cur %= p
		need := (p + cur - sum) % p
		if forePos, ok := pos[need]; ok {
			res = min(res, i-forePos)
		}
		if res == 1 && n != 1 {
			return res
		}
		pos[cur] = i
	}
	if res == n {
		return -1
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
