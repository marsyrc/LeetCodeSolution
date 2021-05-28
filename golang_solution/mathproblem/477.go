package mathproblem

func totalHammingDistance(nums []int) (ans int) {
	n := len(nums)
	for i := 31; i >= 0; i-- {
		cnt := 0
		for _, num := range nums {
			cnt += num >> i & 1
		}
		ans += cnt * (n - cnt)
	}
	return
}
