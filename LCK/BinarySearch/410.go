package binarysearch

import "math"

//转换成找合适的间距问题
func splitArray(nums []int, m int) int {
	low := max(nums...)
	high := sum(nums...)
	for low < high {
		mid := low + (high-low)/2
		cnt := 1
		tmp := 0
		for _, v := range nums {
			tmp += v
			if tmp > mid {
				cnt++
				tmp = v
			}
		}
		//划分的子数组多了，说明单个子数组和mid找小了
		if cnt > m {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}

func max(a ...int) int {
	res := math.MinInt32
	for _, v := range a {
		if v > res {
			res = v
		}
	}
	return res
}

func sum(a ...int) int {
	res := 0
	for _, v := range a {
		res += v
	}
	return res
}
