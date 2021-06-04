package mathproblem

func getLastMoment(n int, left []int, right []int) int {
	res := 0
	for _, pos := range left {
		res = max(res, pos)
	}
	for _, pos := range right {
		res = max(res, n-pos)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
