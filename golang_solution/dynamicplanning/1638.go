package dynamicplanning

func countSubstrings(s string, t string) int {
	m, n := len(s), len(t)
	ans := 0
	for delta := -m + 1; delta < n; delta++ {
		i, j := 0, 0
		if delta <= 0 {
			i = -delta
		} else {
			j = delta
		}

		fij, gij := 0, 0
		for i < m && j < n {
			if s[i] == s[j] {
				gij++
			} else {
				fij = gij + 1
				gij = 0
			}
			ans += fij
			i++
			j++
		}
	}
	return ans
}
