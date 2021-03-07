package twopointers

import "strconv"

func compress(chars []byte) int {
	if len(chars) == 0 {
		return 0
	}
	l, r := 0, 0
	i := 0
	for i < len(chars) && r < len(chars) {
		cur := chars[r]
		chars[i] = cur
		for r < len(chars) && cur == chars[r] {
			r++
		}
		if r-l == 1 {
			i++
			l = r
			continue
		}
		if r >= len(chars) || cur != chars[r] {
			i++
			dis := strconv.Itoa(r - l)
			for j := 0; j < len(dis) && i < len(chars); j++ {
				chars[i] = dis[j]
				i++
			}
		}
		l = r
	}
	return len(chars[:i])
}
