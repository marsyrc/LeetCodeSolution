package stringproblem

func convert(s string, numRows int) string {
	mod := 2*numRows - 2
	if numRows >= len(s) || numRows == 1 {
		return s
	}
	cache := []byte{}
	//ergodic by rows
	for i := 0; i < numRows; i++ {
		if i == 0 || i == numRows-1 {
			for cur := i; cur < len(s); cur += mod {
				cache = append(cache, s[cur])
			}
		} else {
			for cur := i; cur < len(s); cur += mod {
				cache = append(cache, s[cur])
				if next := cur + 2*(numRows-1-cur%mod); next >= len(s) {
					break
				} else {
					cache = append(cache, s[next])
				}
			}
		}
	}
	return string(cache)
}
