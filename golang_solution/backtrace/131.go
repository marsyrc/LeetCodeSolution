package backtrace

func partition(s string) [][]string {
	if len(s) == 0 {
		return [][]string{}
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
			curStr := s[pos : i+1]
			if !isHuiwen(curStr) {
				continue
			}
			path = append(path, curStr)
			backtrace(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtrace(0)
	return res
}

func isHuiwen(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
