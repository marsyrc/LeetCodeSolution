package binarysearch

func minSpeedOnTime(dist []int, hour float64) int {
	l, r := 1, 10000000
	if float64(len(dist))-1 > hour {
		return -1
	}
	canReach := func(speed int) bool {
		var total float64
		for i := 0; i < len(dist)-1; i++ {
			total += float64(dist[i] / speed)
			if dist[i]%speed != 0 {
				total += 1
			}
		}
		total += float64(dist[len(dist)-1]) / float64(speed)
		if total <= hour {
			return true
		}
		return false
	}
	for l < r {
		mid := (l + r) / 2
		if canReach(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
