package dynamicplanning

func minimumDeletions(s string) int {
	n := len(s)
	bCnt := 0
	res := 0
	for i := 0; i < n; i++ {
		if s[i] == 'a' {
			res = min(res+1, bCnt)
		} else {
			bCnt++
		}
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
