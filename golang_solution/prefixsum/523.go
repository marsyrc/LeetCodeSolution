package prefixsum

func checkSubarraySum(nums []int, k int) bool {
	hm := make(map[int]int)
	hm[0] = -1
	cur := 0
	for i := 0; i < len(nums); i++ {
		cur += nums[i]
		cur %= k

		if _, ok := hm[cur]; ok {
			if i-hm[cur] > 1 {
				return true
			}
		} else {
			hm[cur] = i
		}
	}
	return false
}
