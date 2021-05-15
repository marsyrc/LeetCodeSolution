package backtrace

func constructDistancedSequence(n int) []int {
	l := (n-2+1)*2 + 1
	res := make([]int, l)
	cnt := make([]int, n)
	cnt[0] = 1
	path := []int{}
	var bt func(p []int)
	flag := false
	bt = func(p []int) {
		// fmt.Println(p)
		if len(p) == l {
			copy(res, p)
			flag = true
			return
		}
		//make sure that child node is legal as defined
		//
		for i := n; i > 0; i-- {
			if i == 1 { //special case
				if cnt[0] == 1 {
					p = append(p, i)
					cnt[0]++
					bt(p)
					p = p[:len(p)-1]
					cnt[0]--
				}
				continue
			}
			if !flag && cnt[i-1] == 0 { //not used
				p = append(p, i)
				cnt[i-1]++
				bt(p)
				p = p[:len(p)-1]
				cnt[i-1]--
			} else if !flag && cnt[i-1] == 1 {
				if len(p) >= i && p[len(p)-i] == i {
					p = append(p, i)
					cnt[i-1]++
					bt(p)
					p = p[:len(p)-1]
					cnt[i-1]--
				}
			}
		}
	}
	bt(path)
	return res
}
