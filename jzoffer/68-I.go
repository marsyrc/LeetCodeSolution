package jianzhioffer

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先
*/
func lowestCommonAncestor(root, p, q *TreeNode) (ancestor *TreeNode) {
	ancestor = root
	for {
		if p.Val < ancestor.Val && q.Val < ancestor.Val {
			ancestor = ancestor.Left
		} else if p.Val > ancestor.Val && q.Val > ancestor.Val {
			ancestor = ancestor.Right
		} else {
			return
		}
	}
}
