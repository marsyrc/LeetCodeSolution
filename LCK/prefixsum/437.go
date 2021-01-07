package prefixsum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 辅助函数，用于前序遍历二叉树，同时记录下遍历过程中的前进路径和以及出现的次数
// 输入是二叉树，前进路径和，目标值sum，以及辅助哈希表,输出是路径和等于sum的路径数量
func dfs(root *TreeNode, cur, sum int, prefixSum map[int]int) int {
	if root == nil {
		return 0
	}
	cur += root.Val
	cnt := 0
	if v, ok := prefixSum[cur-sum]; ok {
		cnt = v
	}
	// 更新前进路径和cur出现的次数
	prefixSum[cur]++
	// 接下来递归求解左右子树,并把返回的路径数量加到cnt上
	cnt += dfs(root.Left, cur, sum, prefixSum)
	cnt += dfs(root.Right, cur, sum, prefixSum)
	prefixSum[cur]-- // 退递归后需要把现在的前进路径和数量-1
	// 下一步要回溯到递归的上一层进行处理
	return cnt
}

// 保存前进路径和的方法 Time: O(n), Space: O(n)
func pathSum(root *TreeNode, sum int) int {
	prefixSum := make(map[int]int)      // 定义辅助哈希表用于存储每个前进路径和以及它出现的次数
	prefixSum[0] = 1                    // 接着把处理边界的0，1数对加入到哈希表
	return dfs(root, 0, sum, prefixSum) // 最后只要调用辅助函数即可
}
