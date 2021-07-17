package binarysearch

func arrangeCoins(n int) int {
	n *= 2 
	l, r := 0, n 
	for l <= r {
		mid := (l + r) / 2 
		cur := mid * (mid + 1)
		if cur > n {
			r = mid - 1
		} else if cur < n {
            l = mid + 1
		} else {
            return mid 
        }
	}
	return r
}