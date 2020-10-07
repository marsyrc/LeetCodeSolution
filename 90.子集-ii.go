import "sort"

/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 */

// @lc code=start
func subsetsWithDup(nums []int) [][]int {
	// 保存最终结果
	ans := make([][]int, 0)
	//保存中间结果
	list := make([]int, 0)

	sort.Ints(nums)
	backtrack(nums, 0, list, &ans)
	return ans
}

func backtrack(nums []int, pos int, list []int, ans *[][]int) {
	res := make([]int, len(list))
	copy(res, list)
	*ans = append(*ans, res)

	for i := pos; i < len(nums); i++ {
		//如果是重复元素，则跳过
		if i != pos && nums[i] == nums[i-1] {
			continue
		}
		list = append(list, nums[i])
		backtrack(nums, i+1, list, ans)
		list = list[:len(list)-1]
	}
}

// @lc code=end
