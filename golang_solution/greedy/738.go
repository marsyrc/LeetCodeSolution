package greedy

import "strconv"

func monotoneIncreasingDigits(N int) int {
	s := strconv.Itoa(N)
	ss := []byte(s)
	n := len(ss)
	if n <= 1 {
		return N
	}
	lastNine := n
	for i := n - 2; i >= 0; i-- {
		if ss[i] > ss[i+1] {
			ss[i] -= 1
			for j := i + 1; j < lastNine; j++ {
				ss[j] = '9'
			}
			lastNine = i + 1
		}
	}
	str := string(ss)
	res, _ := strconv.Atoi(str)
	return res
}
