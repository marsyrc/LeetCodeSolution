package dynamicplanning

func longestPalindrome(word1 string, word2 string) int {
	return longestPalindromeSubseq(word1+word2, len(word1))
}

func longestPalindromeSubseq(s string, x int) int {
	res := 0
	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		dp[i][i] = 1
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
				if i < x && j >= x {
					res = max(res, dp[i][j])
				}
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
