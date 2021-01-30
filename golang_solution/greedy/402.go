package greedy

import "strings"

//贪心、单调栈
func removeKdigits(num string, k int) string {
	s := []byte{}
	for i := 0; i < len(num); i++ {
		for k > 0 && len(s) != 0 && num[i] < s[len(s)-1] {
			s = s[:len(s)-1]
			k--
		}
		s = append(s, num[i])
	}
	s = s[:len(s)-k]
	res := strings.TrimLeft(string(s), "0")
	if res == "" {
		return "0"
	}
	return res
}
