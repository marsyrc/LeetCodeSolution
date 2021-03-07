/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */

// @lc code=start
func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}

	//从后向前查找第一个相邻升序的元素对
	i, j, k := len(nums)-2, len(nums)-1, len(nums)-1
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}

	//在[j,end)中从后开始查找第一个满足nums[k]大于nums[i]的k.并交换
	if i >= 0 {
		for nums[i] >= nums[k] {
			k--
		}
		nums[i], nums[k] = nums[k], nums[i]
	}

	//将[j,end)逆序
	for m, n := j, len(nums)-1; m < n; m, n = m+1, n-1 {
		nums[m], nums[n] = nums[n], nums[m]
	}
}

// @lc code=end
