package binarysearch

func singleNonDuplicate(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)/2
		if mid%2 == 1 {
			mid--
		}
		if nums[mid+1] == nums[mid] {
			l = mid + 2
		} else {
			r = mid
		}
	}
	return nums[l]
}
