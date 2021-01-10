package binarysearchtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	list := []*TreeNode{}
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		s := []*TreeNode{}
		for len(s) > 0 || root != nil {
			for root != nil {
				s = append(s, root)
				root = root.Left
			}
			root = s[len(s)-1]
			s = s[:len(s)-1]
			list = append(list, root)
			root = root.Right
		}
	}
	inorder(root)
	return list[k-1].Val
}
