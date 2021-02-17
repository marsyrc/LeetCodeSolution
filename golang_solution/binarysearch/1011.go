package binarysearch

func shipWithinDays(weights []int, D int) int {
	canFinish := func(cap int) bool {
		i := 0
		cnt := 0
		cur := 0
		for i < len(weights) {
			for i < len(weights) && cur+weights[i] <= cap {
				cur += weights[i]
				i++
			}
			cnt++
			cur = 0
		}
		return cnt <= D
	}
	l, r := -1, 0
	for _, v := range weights {
		if v > l {
			l = v
		}
		r += v
	}
	for l < r {
		mid := l + (r-l)/2
		if canFinish(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
