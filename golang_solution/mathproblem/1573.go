package mathproblem

func numWays(s string) int {
	gap1, gap2 := 0, 0
	cnt := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			cnt++
		}
	}

	if cnt%3 != 0 {
		return 0
	}

	if cnt == 0 {
		return (len(s) - 1) * (len(s) - 2) / 2 % (1e9 + 7)
	}

	cnt /= 3
	n := 0
	i := 0

	for i < len(s) && n < cnt {
		if i < len(s) && s[i] == '1' {
			n++
		}
		i++
	}

	for i < len(s) && s[i] == '0' {
		gap1++
		i++
	}
	n = 0

	for i < len(s) && n < cnt {
		if s[i] == '1' {
			n++
		}
		i++
	}

	for i < len(s) && s[i] == '0' {
		i++
		gap2++
	}

	return ((gap1 + 1) * (gap2 + 1)) % (1e9 + 7)
}
