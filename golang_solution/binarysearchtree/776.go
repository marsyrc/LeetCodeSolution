package binarysearchtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func splitBST(root *TreeNode, V int) []*TreeNode {
	if root == nil {
		return []*TreeNode{nil, nil}
	}
	//root属于结果的左子树还是右子树
	if root.Val <= V {
		ans := splitBST(root.Right, V)
		root.Right = ans[0]
		ans[0] = root
		return ans
	} else {
		ans := splitBST(root.Left, V)
		root.Left = ans[1]
		ans[1] = root
		return ans
	}
}
