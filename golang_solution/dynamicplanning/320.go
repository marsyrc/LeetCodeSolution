package dynamicplanning

func minPatches(nums []int, n int) (patches int) {
	x := 1
	i := 0
	for x <= n {
		if i < len(nums) && nums[i] <= x {
			x += nums[i]
			i++
		} else {
			x *= 2
			patches++
		}
	}
	return
}
