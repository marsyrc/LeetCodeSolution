package hashmap

func numPairsDivisibleBy60(time []int) (res int) {
	m := map[int]int{}
	for i := 0; i < len(time); i++ {
		cur := time[i] % 60
		if cur == 0 {
			res += m[0]
			m[0]++
			continue
		}
		if cnt, ok := m[60-cur]; ok {
			res += cnt
		}
		m[cur]++
	}
	return
}
