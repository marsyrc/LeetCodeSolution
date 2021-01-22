package binarysearch

import "sort"

//用二分搜索一个尽可能大的间隔
func maxDistance(position []int, m int) (ans int) {
	sort.Ints(position)
	low, high := 1, (position[len(position)-1]-position[0])/(m-1)
	for low <= high {
		mid := low + (high-low)/2
		if isLegal(position, m, mid) {
			ans = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return ans
}

func isLegal(position []int, m int, mid int) bool {
	cnt := 1
	slow, fast := 0, 1
	for fast < len(position) {
		for fast < len(position) && position[fast]-position[slow] < mid {
			fast++
		}
		if fast < len(position) {
			cnt++
		} else {
			break
		}
		slow = fast
		fast++
	}
	if cnt >= m {
		return true
	}
	return false
}
