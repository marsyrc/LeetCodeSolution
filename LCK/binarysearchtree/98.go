package binarysearchtree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	pre := math.MinInt64
	var inorder func(root *TreeNode) bool
	inorder = func(root *TreeNode) bool {
		if root == nil {
			return true
		}

		if !inorder(root.Left) {
			return false
		}

		if pre >= root.Val {
			return false
		}

		pre = root.Val

		return inorder(root.Right)
	}
	return inorder(root)
}
