package dynamicplanning

import "sort"

//LongestIncreasingSequence DP problem
//russia env
func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		return envelopes[i][0] < envelopes[j][0] || (envelopes[i][0] == envelopes[j][0] && envelopes[i][1] > envelopes[j][1])
	})
	dp := make([]int, len(envelopes))
	if len(envelopes) == 0 {
		return 0
	}
	dp[0] = 1
	maxNum := 1
	for i := 1; i < len(envelopes); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if envelopes[i][1] > envelopes[j][1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxNum = max(maxNum, dp[i])
	}
	return maxNum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
