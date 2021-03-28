package mathproblem

func canArrange(arr []int, k int) bool {
	cnt := make([]int, k)
	for _, a := range arr {
		cnt[(a%k+k)%k]++
	}

	for i := 1; i+i < k; i++ {
		if cnt[i] != cnt[k-i] {
			return false
		}
	}
	if cnt[0]&1 != 0 {
		return false
	}
	return true
}
