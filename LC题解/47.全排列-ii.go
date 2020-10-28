import "sort"

/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
	var ans [][]int
	var path []int

	sort.Ints(nums)
	n := len(nums)

	vis := make([]bool, n)
	var backTrack func(int)
	backTrack = func(index int) {
		if index == n {
			ans = append(ans, append([]int(nil), path...))
			return
		}

		for i, v := range nums {
			if vis[i] || i > 0 && !vis[i-1] && v == nums[i-1] {
				continue
			}
			path = append(path, v)
			vis[i] = true

			backTrack(index + 1)
			vis[i] = false
			path = path[:len(path)-1]

		}
	}
	backTrack(0)
	return ans
}

// @lc code=end
