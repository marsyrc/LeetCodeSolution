package hashmap

func findTheLongestSubstring(s string) int {
	hm := make(map[int]int)
	res := 0
	state := 0
	hm[0] = 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'a':
			state ^= 1
		case 'e':
			state ^= 2
		case 'i':
			state ^= 4
		case 'o':
			state ^= 8
		case 'u':
			state ^= 16
		}
		if last, ok := hm[state]; ok {
			res = max(res, i-last+1)
		} else {
			hm[state] = i + 1
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
