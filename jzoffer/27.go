package jianzhioffer

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	l := mirrorTree(root.Left)
	r := mirrorTree(root.Right)
	root.Right = l
	root.Left = r
	return root
}
