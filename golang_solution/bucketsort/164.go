package bucketsort

import "math"

//在桶内只存最大最小值，判断桶间距离
type pair struct {
	max int
	min int
}

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	minVal := min(nums...)
	maxVal := max(nums...)
	n := len(nums)
	step := max((maxVal-minVal)/(n-1), 1)
	bucketsize := (maxVal-minVal)/step + 1

	buckets := make([]pair, bucketsize)
	for i := range buckets {
		buckets[i] = pair{-1, -1}
	}

	//index nums[i] into buckets
	for _, v := range nums {
		index := (v - minVal) / step
		if buckets[index].max == -1 {
			buckets[index].max = v
			buckets[index].min = v
		} else {
			buckets[index].max = max(buckets[index].max, v)
			buckets[index].min = min(buckets[index].min, v)
		}
	}
	ans := 0
	prev := -1
	for i, b := range buckets {
		if b.min == -1 {
			continue
		}
		if prev != -1 {
			ans = max(ans, b.min-buckets[prev].max)
		}
		prev = i
	}
	return ans
}

func min(a ...int) int {
	res := math.MaxInt32
	for _, v := range a {
		if v < res {
			res = v
		}
	}
	return res
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
