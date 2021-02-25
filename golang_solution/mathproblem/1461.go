package mathproblem

import "math"

//滑动窗口+位运算
func hasAllCodes(s string, k int) bool {
	m := make(map[int]int)
	cur := 0
	i := 0
	for i < len(s) && i < k {
		cur <<= 1
		cur += int(s[i] - '0')
		i++
	}
	if i == len(s) {
		return false
	}
	m[cur]++
	mask := int(math.Pow(2, float64(k))) - 1
	for i := k; i < len(s); i++ {
		cur <<= 1
		cur &= mask
		cur += int(s[i] - '0')
		m[cur]++
	}
	return len(m) == int(math.Pow(2, float64(k)))
}
