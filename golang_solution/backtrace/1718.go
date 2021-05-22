package backtrace

func constructDistancedSequence(n int) []int {
	l := 2*n - 1
	res := make([]int, l)
	cnt := make([]int, n)
	var bt func(idx int) bool
	bt = func(idx int) bool {
		if idx == l {
			return true
		}
		if res[idx] != 0 {
			return bt(idx + 1)
		}
		for i := n; i > 1; i-- {
			if cnt[i-1] == 0 && i+idx < l && res[i+idx] == 0 { //not used
				res[idx] = i
				res[idx+i] = i
				cnt[i-1] = 1

				if bt(idx + 1) {
					return true
				}

				res[idx] = 0
				res[idx+i] = 0
				cnt[i-1] = 0
			}
		}
		if cnt[0] == 0 {
			res[idx] = 1
			cnt[0] = 1
			if bt(idx + 1) {
				return true
			}
			cnt[0] = 0
			res[idx] = 0
		}
		return false
	}
	bt(0)
	return res
}
