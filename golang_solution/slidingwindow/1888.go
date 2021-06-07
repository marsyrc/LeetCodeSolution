package slidingwindow

import "math"

func minFlips(s string) int {
	l := s + s
	n := len(s)
	cnt := 0
	for i := range s {
		if i%2 == 0 {
			if s[i] != '1' {
				cnt++
			}
		} else {
			if s[i] != '0' {
				cnt++
			}
		}
	}
	target := "10"
	res := min(cnt, n-cnt)
	for i := 0; i < n; i++ {
		if l[i] != target[i%2] {
			cnt--
		}
		if l[i+n] != target[(i+n)%2] {
			cnt++
		}
		res = min(res, cnt, n-cnt)
	}
	return res
}

func min(a ...int) int {
	res := math.MaxInt32
	for _, v := range a {
		if v < res {
			res = v
		}
	}
	return res
}
