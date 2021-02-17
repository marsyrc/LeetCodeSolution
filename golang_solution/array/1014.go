package array

func maxScoreSightseeingPair(A []int) int {
	mx := A[0] + 0
	ans := 0
	for i := 1; i < len(A); i++ {
		ans = max(A[i]-i+mx, ans)
		mx = max(mx, A[i]+i)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
