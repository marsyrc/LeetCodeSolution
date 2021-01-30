package dynamicplanning

func numberOfArithmeticSlices(A []int) int {
	n := len(A)
	dp := make([]int, n)
	sum := 0
	for i := 2; i < n; i++ {
		if A[i]-A[i-1] == A[i-1]-A[i-2] {
			dp[i] = dp[i-1] + 1
			sum += dp[i]
		}
	}
	return sum
}
