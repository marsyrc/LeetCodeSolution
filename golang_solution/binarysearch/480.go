package binarysearch

import "sort"

//求中位数，用二分法查找插入位置
func medianSlidingWindow(nums []int, k int) []float64 {
	window := make([]int, k)
	copy(window, nums[:k])
	sort.Ints(window)
	i := 0
	j := k
	res := make([]float64, len(nums)-k+1)
	res[0] = calMedian(window)
	for j < len(nums) {
		deleteIdx := bisect(window, nums[i])
		window = append(window[:deleteIdx], window[deleteIdx+1:]...)
		addIdx := bisect(window, nums[j])
		tails := append([]int{nums[j]}, window[addIdx:]...)
		window = append(window[:addIdx], tails...)
		j++
		i++
		res[i] = calMedian(window)
	}
	return res
}

func bisect(nums []int, val int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == val {
			return mid
		} else if nums[mid] > val {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}

func calMedian(nums []int) float64 {
	n := len(nums)
	if n%2 == 0 {
		return float64(nums[n/2])/float64(2) + float64(nums[n/2-1])/float64(2)
	} else {
		return float64(nums[n/2])
	}
}
