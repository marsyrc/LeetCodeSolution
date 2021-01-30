package greedy

func minPatches(nums []int, n int) int {
	res := 0
	x := 1
	i := 0
	for x <= n {
		if i < len(nums) && nums[i] <= x {
			x += nums[i]
			i++
		} else {
			res++
			x *= 2
		}
	}
	return res
}
