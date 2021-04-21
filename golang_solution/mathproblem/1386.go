package mathproblem

func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
	M := map[int]int{}
	for _, v := range reservedSeats {
		r, c := v[0], v[1]
		if c == 1 || c == 10 {
			continue
		}
		M[r] += 1 << (9 - c)
	}
	cnt := n * 2
	for _, v := range M {
		if v&0xF0 == 0 || v&0x3C == 0 || v&0x0F == 0 {
			cnt--
		} else {
			cnt -= 2
		}
	}

	return cnt
}
