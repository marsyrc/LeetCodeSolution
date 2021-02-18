package jianzhioffer

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargest(root *TreeNode, k int) int {
	list := []int{}
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Right)
		list = append(list, node.Val)
		inorder(node.Left)
	}
	inorder(root)
	return list[k-1]
}
