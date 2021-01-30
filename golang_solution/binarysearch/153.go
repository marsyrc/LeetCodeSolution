package binarysearch

func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	n := len(nums)
	for l <= r {
		mid := l + (r-l)/2
		if mid > 0 && nums[mid] < nums[mid-1] {
			return nums[mid]
		}
		if mid < n-1 && nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}
		if nums[mid] > nums[l] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return nums[0]
}
