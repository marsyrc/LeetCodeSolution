package binarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	s := []*TreeNode{}
	for root != nil || len(s) > 0 {
		for root != nil {
			res = append(res, root.Val)
			s = append(s, root)
			root = root.Left
		}
		root = s[len(s)-1]
		s = s[:len(s)-1]
		root = root.Right
	}
	return res
}
