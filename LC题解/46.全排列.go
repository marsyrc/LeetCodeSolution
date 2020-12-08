/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
	var path []int
	var used = make([]bool, len(nums))
	result = [][]int{}
	backTrack(nums, path, used)
	return result
}
func backTrack(nums, path []int, used []bool) {
	if len(nums) == len(path) {

		//深拷贝
		tmp := make([]int, len(nums))
		copy(tmp, path)
		//满足终止条件，加入结果
		result = append(result, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if !used[i] {
			used[i] = true
			//做选择
			path = append(path, nums[i])
			//递归
			backTrack(nums, path, used)
			//撤销
			path = path[:len(path)-1]
			used[i] = false
		}
	}
}

// @lc code=end
