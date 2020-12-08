import "math"

// 152. 乘积最大子数组

// 给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。

// 示例 1:

// 输入: [2,3,-2,4]
// 输出: 6
// 解释: 子数组 [2,3] 有最大乘积 6。

// 示例 2:

// 输入: [-2,0,-1]
// 输出: 0
// 解释: 结果不能为 2, 因为 [-2,-1] 不是子数组

func maxProduct(nums []int) int {
	n := len(nums)
	maxDP := make([]int, n)
	minDP := make([]int, n)

	if n == 0 {
		return 0
	}
	maxDP[0], minDP[0] = nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		maxDP[i] = max(max(maxDP[i-1]*nums[i], minDP[i-1]*nums[i]), nums[i])
		minDP[i] = min(min(maxDP[i-1]*nums[i], minDP[i-1]*nums[i]), nums[i])
	}
	ans := math.MinInt64
	for _, v := range maxDP {
		if v > ans {
			ans = v
		}
	}
	return ans
}

/*
压缩后
*/
/*
func maxProduct(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

    maxDP, minDP := nums[0], nums[0]
    ans := nums[0]
	for i := 1; i < len(nums); i++ {
        tmp := maxDP
		maxDP = max(max(maxDP * nums[i], minDP * nums[i]), nums[i])
		minDP = min(min(tmp * nums[i], minDP * nums[i]), nums[i])
        ans = max(ans, maxDP)
	}

    return ans
}
*/
