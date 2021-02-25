package prefixsum

func minOperations(boxes string) []int {
	dp := make([]int, len(boxes))
	right := 0
	for i, v := range boxes {
		if v == '1' {
			dp[0] += i
			right++
		}
	}
	left := 0
	for i := 1; i < len(boxes); i++ {
		if boxes[i-1] == '1' {
			left++
			right--
		}
		dp[i] = dp[i-1] - right + left
	}
	return dp
}
