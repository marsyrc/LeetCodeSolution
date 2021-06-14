package binarysearch

func maximumRemovals(s string, p string, removable []int) int {
	check := func(ss string) bool {
		i := 0
		for _, b := range ss {
			if byte(b) == p[i] {
				if i++; i == len(p) {
					return true
				}
			}
		}
		return false
	}

	ans := 0
	l, r := 0, len(removable)
	for l <= r {
		mid := (l + r) / 2

		del := make([]bool, len(s))
		for i := 0; i < mid; i++ {
			del[removable[i]] = true
		}
		temp := []byte{}
		for i := range s {
			if !del[i] {
				temp = append(temp, s[i])
			}
		}

		if check(string(temp)) {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return ans
}
