package backtrace

func maxUniqueSplit(s string) int {
	res := 1
	var dfs func(str string)
	memo := make(map[string]struct{})
	dfs = func(str string) {
		if len(memo)+len(str) < res {
			return
		}
		if len(str) == 0 {
			res = max(res, len(memo))
		}

		for i := 0; i < len(str); i++ {
			cur := str[:i+1]
			if _, ok := memo[cur]; !ok {
				memo[cur] = struct{}{}

				dfs(str[i+1:])

				delete(memo, cur)
			}
		}
	}
	dfs(s)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
