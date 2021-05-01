package greedy

func canConstruct(s string, k int) bool {
	n := len(s)
	cnt := [26]int{}
	for i := 0; i < n; i++ {
		cnt[int(s[i]-'a')]++
	}
	oddCnt := 0
	for i := 0; i < 26; i++ {
		if cnt[i]&1 != 0 {
			oddCnt++
		}
	}
	if n < k || oddCnt > k {
		return false
	}
	return true
}
