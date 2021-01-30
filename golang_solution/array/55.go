package array

func canJump(nums []int) bool {
	maxPos := nums[0] + 0
	for i := 0; i < len(nums); i++ {
		if i > maxPos {
			return false
		}
		maxPos = max(maxPos, nums[i]+i)
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
