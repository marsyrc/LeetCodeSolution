package prefixsum

func wonderfulSubstrings(word string) (res int64) {
	cur := 0
	cntMap := make(map[int]int64)
	cntMap[0] = 1
	for i := range word {
		cur ^= (1 << int(word[i]-'a'))
		if allEven, ok := cntMap[cur]; ok {
			res += allEven
		}
		for j := 0; j < 10; j++ {
			if sameCnt, ok := cntMap[cur^(1<<j)]; ok {
				res += sameCnt
			}
		}
		cntMap[cur]++
	}
	return
}
