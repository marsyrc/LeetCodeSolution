package dynamicplanning

func canReach(s string, minJump int, maxJump int) bool {
	n := len(s)
	dp := make([]bool, n)
	cnt := 1
	dp[0] = true

	for i := minJump; i < n; i++ {
		if s[i] == '0' && cnt > 0 {
			dp[i] = true
		}

		if i-maxJump >= 0 && dp[i-maxJump] {
			cnt--
		}

		if dp[i-minJump+1] {
			cnt++
		}
	}
	return dp[n-1]
}
