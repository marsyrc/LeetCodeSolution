package slidingwindow

import "math"

func maxVowels(s string, k int) int {
	cnt := 0
	i := 0
	for i < len(s) && i < k {
		if isVowel(s[i]) {
			cnt++
		}
		i++
	}
	if i == len(s) {
		return cnt
	}
	i = k - 1
	res := cnt
	for i < len(s)-1 {
		i++
		if isVowel(s[i]) {
			cnt++
		}
		if isVowel(s[i-k]) {
			cnt--
		}
		res = max(res, cnt)
	}
	return res
}

func isVowel(b byte) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u'
}

func max(a ...int) int {
	res := math.MinInt32
	for _, v := range a {
		if v > res {
			res = v
		}
	}
	return res
}
