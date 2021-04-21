package greedy

func partitionDisjoint(A []int) int {
	maxVal := A[0]
	leftMax := A[0]
	pos := 0
	for i := 0; i < len(A); i++ {
		maxVal = max(maxVal, A[i])
		if A[i] >= leftMax {
			continue
		}
		leftMax = maxVal
		pos = i
	}
	return pos + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
