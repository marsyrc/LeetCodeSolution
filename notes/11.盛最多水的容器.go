/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	ans := 0
	for left < right {
		currentArea := (right - left) * min(height[right], height[left])
		ans = max(currentArea, ans)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end
