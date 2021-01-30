package binarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := countLevel(root.Left)
	right := countLevel(root.Right)
	if right == left {
		return countNodes(root.Right) + (1 << left) - 1 + 1
	} else {
		return countNodes(root.Left) + (1 << right) - 1 + 1
	}
}

func countLevel(root *TreeNode) int {
	level := 0
	for root != nil {
		level++
		root = root.Left
	}
	return level
}
