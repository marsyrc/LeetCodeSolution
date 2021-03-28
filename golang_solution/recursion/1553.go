package recursion

func minDays(n int) int {
	memory := make(map[int]int)
	memory[0] = 0
	memory[1] = 1

	var helper func(n int) int
	helper = func(n int) int {
		if days, ok := memory[n]; ok {
			return days
		}
		memory[n] = min(helper(n/2)+n%2+1, helper(n/3)+n%3+1)
		return memory[n]
	}
	return helper(n)
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
