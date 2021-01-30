package stringproblem

import "sort"

func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	sort.Strings(strs)
	l := strs[0]
	r := strs[len(strs)-1]

	ans := ""
	i := 0
	for i < min(len(l), len(r)) && l[i] == r[i] {
		ans += string(l[i])
		i++
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
