package array

//floyd判环算法
func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	slow = nums[slow]
	fast = nums[nums[fast]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	//slow回起点
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return fast
}
