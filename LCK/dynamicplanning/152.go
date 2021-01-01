package dynamicplanning

import "math"

func maxProduct(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	ans := nums[0]
	productMax, productMin := nums[0], nums[0]
	for i := 1; i < n; i++ {
		productMax, productMin = max(productMax*nums[i], productMin*nums[i], nums[i]), min(productMax*nums[i], productMin*nums[i], nums[i])
		ans = max(ans, productMax)
	}
	ans = max(ans, productMax)
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
