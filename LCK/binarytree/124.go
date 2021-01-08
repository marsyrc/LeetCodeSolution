package binarytree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	res := math.MinInt32
	var nodeSum func(node *TreeNode) int
	nodeSum = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftSum := max(nodeSum(node.Left), 0)
		rightSum := max(nodeSum(node.Right), 0)

		curRes := leftSum + rightSum + node.Val
		res = max(res, curRes)
		return node.Val + max(leftSum, rightSum)
	}
	nodeSum(root)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
