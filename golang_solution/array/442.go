package array

//in-place hash
func findDuplicates(nums []int) []int {
	ans := []int{}
	for i := 0; i < len(nums); i++ {
		idx := abs(nums[i]) - 1
		if nums[idx] < 0 {
			ans = append(ans, idx+1)
		}
		nums[idx] = -nums[idx]
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
