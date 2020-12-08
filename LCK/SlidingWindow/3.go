package SlidingWindow

func lengthOfLongestSubstring(s string) int {
	window := make(map[byte]int)
	if len(s) == 0 {
		return 0
	}
	left, right := 0, 0
	res := math.MinInt32
	ss := []byte(s)
	for right < len(s) {
		c := ss[right]
		right++ 

		window[c]++

		for window[c] != 1 {
			d := ss[left]
			window[d]--
			left++
		}

		res = max(res, right + 1 - left)
	}
	return res
}

func max (a ...int) int {
	res := math.MinInt32
	for _, v := range a {
		if v > res {
			res = v
		}
	}
	return res
}