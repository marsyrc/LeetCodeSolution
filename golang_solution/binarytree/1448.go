package binarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) (ans int) {
	var dfs func(node *TreeNode, maxValue int)
	dfs = func(node *TreeNode, maxValue int) {
		if node == nil {
			return
		}
		if node.Val >= maxValue {
			ans++
		}
		dfs(node.Left, max(maxValue, node.Val))
		dfs(node.Right, max(maxValue, node.Val))
	}
	dfs(root, -10001)
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
