package stack

//维护一个单调递增栈
func largestRectangleArea(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}
	res := 0
	s := []int{}
	for i := 0; i < n; i++ {
		for len(s) > 0 && heights[s[len(s)-1]] > heights[i] {
			length := heights[s[len(s)-1]]
			s = s[:len(s)-1]
			weight := i
			if len(s) != 0 {
				weight = i - s[len(s)-1] - 1
			}
			res = max(res, length*weight)
		}
		s = append(s, i)
	}

	for len(s) > 0 {
		length := heights[s[len(s)-1]]
		s = s[:len(s)-1]
		weight := n
		if len(s) != 0 {
			weight = n - s[len(s)-1] - 1
		}
		res = max(res, length*weight)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
