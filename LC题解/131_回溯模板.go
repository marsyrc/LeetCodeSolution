/**
131. 分割回文串

给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。

返回 s 所有可能的分割方案。

示例:

输入: "aab"
输出:
[
  ["aa","b"],
  ["a","a","b"]
]

**/
func partition(s string) [][]string {
	if len(s) == 0 {
		return [][]string{}
	}
	var res [][]string
	var backtrace func(s string, path []string, pos int)
	backtrace = func(s string, path []string, pos int) {
		if pos == len(s) {
			tmp := make([]string, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for i := pos; i < len(s); i++ {
			if !isHuiwen(s[pos:i+1]) {
				continue
			}
			path = append(path, s[pos:i+1])
			backtrace(s, path, i + 1)
			path = path[:len(path)-1]
		}
	}
	backtrace(s, []string{}, 0)
	return res
}

func isHuiwen(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		} else {
			left++
			right--
		}
	}
	return true
}

