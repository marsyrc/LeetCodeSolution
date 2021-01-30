package binarysearch

func findMin(nums []int) int {
	n := len(nums)
	l, r := 0, n-1
	if nums[l] < nums[r] || n == 1 {
		return nums[0]
	}
	for l < r {
		if l > 0 && nums[l] < nums[l-1] {
			return nums[l]
		}
		mid := l + (r-l)/2
		if nums[mid] > nums[l] {
			l = mid
		} else if nums[mid] < nums[l] {
			r = mid
		} else {
			l++
		}
	}
	return nums[l]
}
