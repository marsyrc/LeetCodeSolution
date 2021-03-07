//把同一深度的节点弄到一起
func findLeaves(root *TreeNode) [][]int {
	res := [][]int{}
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return -1
		}
		left, right := dfs(node.Left), dfs(node.Right)
		cur := max(left, right) + 1 //当前深度
		if cur >= len(res) {
			res = append(res, []int{})
		}
		res[cur] = append(res[cur], node.Val)
		return cur
	}
	dfs(root)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}