package greedy

func largestMerge(s string, t string) string {
	i, j := 0, 0
	res := ""
	for i < len(s) && j < len(t) {
		if s[i:] > t[j:] {
			res += string(s[i])
			i++
		} else {
			res += string(t[j])
			j++
		}
	}
	res += s[i:] + t[j:]
	return res
}
