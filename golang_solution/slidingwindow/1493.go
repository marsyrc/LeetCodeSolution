package slidingwindow

func longestSubarray(nums []int) int {
	i, j := 0, 0
	res := 0
	cnt := 0

	for j < len(nums) {
		if nums[j] == 0 {
			cnt++
		}
		for cnt > 1 {
			if nums[i] == 0 {
				cnt--
			}
			i++
		}
		res = max(res, j-i+1-1)
		j++
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
