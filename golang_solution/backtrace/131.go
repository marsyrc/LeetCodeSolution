package backtrace

func partition(s string) [][]string {
	if len(s) == 0 {
		return [][]string{}
	}
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}
	for l := 2; l <= n; l++ {
		for i := 0; i+l-1 < n; i++ {
			j := i + l - 1
			if s[i] == s[j] {
				if l == 2 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
		}
	}

	var res [][]string
	var path []string
	var backtrace func(pos int)
	backtrace = func(pos int) {
		if pos == len(s) {
			res = append(res, append([]string{}, path...))
			return
		}
		for i := pos; i < len(s); i++ {
			if !dp[pos][i] {
				continue
			}
			path = append(path, s[pos:i+1])
			backtrace(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtrace(0)
	return res
}
