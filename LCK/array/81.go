package array

//旋转数组，带重复的
func search(nums []int, target int) bool {
	left := 0
	right := len(nums) - 1

	for left <= right {
		//去重复
		for left < right && nums[left] == nums[left+1] {
			left++
		}
		for left < right && nums[right] == nums[right-1] {
			right--
		}
		mid := left + (right-left)/2

		if nums[mid] == target {
			return true
		}
		if nums[mid] <= nums[right] { //mid落在右边
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if target >= nums[0] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return false
}
