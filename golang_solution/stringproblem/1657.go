package stringproblem

import "sort"

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	n := len(word1)
	c1, c2 := make([]int, 26), make([]int, 26)
	for i := 0; i < n; i++ {
		c1[int(word1[i]-'a')]++
		c2[int(word2[i]-'a')]++
	}
	for i := 0; i < 26; i++ {
		if (c1[i] == 0 && c2[i] != 0) || (c1[i] != 0 && c2[i] == 0) {
			return false
		}
	}
	sort.Ints(c1)
	sort.Ints(c2)

	for i := 0; i < 26; i++ {
		if c1[i] != c2[i] {
			return false
		}
	}
	return true
}
