package binarysearch

func longestRepeatingSubstring(s string) int {
	l, r := 0, len(s)-1
	res := 0
	check := func(ll int) bool {
		hm := make(map[string]struct{})
		for i := 0; i+ll-1 < len(s); i++ {
			sub := s[i : i+ll]
			if _, ok := hm[sub]; !ok {
				hm[sub] = struct{}{}
			} else {
				return true
			}
		}
		return false
	}
	for l <= r {
		mid := (l + r) / 2
		if check(mid) {
			res = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return res
}
