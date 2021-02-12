package slidingwindow

func checkInclusion(s1 string, s2 string) bool {
	window := make(map[uint8]int)
	need := make(map[uint8]int)
	for i := 0; i < len(s1); i++ {
		need[s1[i]]++
	}
	left, right := 0, 0
	valid := 0
	for right < len(s2) {
		c := s2[right]
		right++
		if need[c] > 0 {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		for right-left+1 > len(s1) {
			if valid == len(need) {
				return true
			}
			d := s2[left]
			left++
			if need[d] > 0 {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return false
}
