package binarysearch

import "sort"

//二分+双指针
func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	low, high := 0, nums[len(nums)-1]-nums[0]
	for low < high {
		cnt := 0
		mid := low + (high-low)/2
		//双指针计算间距小于mid的对数cnt
		l := 0
		r := 0
		for r < len(nums) {
			for nums[r]-nums[l] > mid {
				l++
			}
			cnt += r - l
			r++
		}
		//cnt多了,说明mid大了,缩小mid
		if cnt >= k {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}
