package twopointers

func numberOfSubstrings(s string) int {
    cnt := [3]int{}
	n := len(s)
	ans := 0
	l, r := 0, -1
	for l < n {
		for r < n && !(cnt[0] > 0 && cnt[1] > 0 && cnt[2] > 0) {
			r++
			if r >= n {
				break
			}
			cnt[s[r] - 'a']++
		}
		ans += n - r
		cnt[s[l] - 'a']--
		l++
	}
	return ans
}