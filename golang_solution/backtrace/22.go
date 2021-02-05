package backtrace

func generateParenthesis(n int) []string {
	path := []byte{}
	ans := []string{}
	var dfs func(left, right int)
	dfs = func(left, right int) {
		if left > n || right > left {
			return
		}
		if left == n && right == n {
			ans = append(ans, string(append([]byte(nil), path...)))
			return
		}
		path = append(path, '(')
		dfs(left+1, right)
		path = path[:len(path)-1]

		path = append(path, ')')
		dfs(left, right+1)
		path = path[:len(path)-1]
	}
	dfs(0, 0)
	return ans
}
