package slidingwindow

func longestOnes(A []int, K int) int {
	l, r := 0, 0
	zeroCnt := 0
	res := 0
	for r < len(A) {
		if A[r] == 0 {
			zeroCnt++
		}
		for zeroCnt > K {
			if A[l] == 0 {
				zeroCnt--
			}
			l++
		}
		res = max(res, r-l+1)
		r++
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
