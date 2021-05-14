package dynamicplanning

import "math"

func maxSumAfterOperation(nums []int) int {
	dp_1, dp_2 := nums[0], nums[0]*nums[0]
	maxV := dp_2
	for i := 1; i < len(nums); i++ {
		t1, t2 := dp_1, dp_2
		dp_1 = max(nums[i], t1+nums[i])
		dp_2 = max(nums[i]*nums[i], t1+nums[i]*nums[i], t2+nums[i])
		maxV = max(maxV, dp_2, dp_1)
	}
	return maxV
}

func max(a ...int) int {
	res := math.MinInt64
	for _, v := range a {
		if v > res {
			res = v
		}
	}
	return res
}
