// 301. 删除无效的括号

// 删除最小数量的无效括号，使得输入的字符串有效，返回所有可能的结果。

// 说明: 输入可能包含了除 ( 和 ) 以外的字符。

// 示例 1:

// 输入: "()())()"
// 输出: ["()()()", "(())()"]

// 示例 2:

// 输入: "(a)())()"
// 输出: ["(a)()()", "(a())()"]

// 示例 3:

// 输入: ")("
// 输出: [""]

func removeInvalidParentheses(s string) []string {
	var isLegal func(s string) bool
	isLegal = func(s string) bool {
		ss := []rune(s)
		first, second := 0, 0
		for _, v := range ss {
			if v == '(' {
				first++
			}
			if v == ')' {
				second++
			}
			if second > first {
				return false
			}
		}
		return second == first
	}

	//bfs
	var ans []string
	level := map[string]bool{
		s: true,
	}
	for len(level) != 0 {
		for v := range level {
			if isLegal(v) {
				ans = append(ans, v)
			}
		}
		if len(ans) != 0 {
			break
		}
		nextLevel := map[string]bool{}
		for v := range level {
			for i := 0; i < len(v); i++ {
				if v[i] == '(' || v[i] == ')' {
					cur := []byte{}
					cur = append(cur, []byte(v)[:i]...)
					cur = append(cur, []byte(v)[i+1:]...)
					nextLevel[string(cur)] = true
				}
			}
		}
		level = nextLevel
	}
	return ans
}