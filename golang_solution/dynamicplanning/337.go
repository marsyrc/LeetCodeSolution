package dynamicplanning

import "math"

//打家劫舍三
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	var dfs func(node *TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		l0, l1 := dfs(node.Left)
		r0, r1 := dfs(node.Right)
		cur0 := max(l0, l1) + max(r0, r1)
		cur1 := node.Val + l0 + r0
		return cur0, cur1
	}
	return max(dfs(root))
}

func max(a ...int) int {
	res := math.MinInt32
	for _, v := range a {
		if res < v {
			res = v
		}
	}
	return res
}
