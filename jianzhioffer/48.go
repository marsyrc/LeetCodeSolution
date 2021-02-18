package jianzhioffer

//slidingwindow
func lengthOfLongestSubstring(s string) int {
	window := make(map[uint8]int)
	left, right := 0, 0
	ans := 0
	for right < len(s) {
		c := s[right]
		right++

		window[c]++

		for window[c] > 1 {
			d := s[left]
			left++

			window[d]--
		}
		ans = max(ans, right-left)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
