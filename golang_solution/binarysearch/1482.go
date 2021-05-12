package binarysearch

func minDays(bloomDay []int, m int, k int) int {
	l, r := 0, 0
	for _, bd := range bloomDay {
		r = max(r, bd)
	}

	if m*k > len(bloomDay) {
		return -1
	}

	check := func(t int) bool {
		cur := 0
		sum := 0
		for i := 0; i < len(bloomDay) && sum < m; i++ {
			if bloomDay[i] <= t {
				cur++
				if cur == k {
					sum++
					cur = 0
				}
			} else {
				cur = 0
			}
		}
		return sum >= m
	}

	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
