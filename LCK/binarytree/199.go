package binarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) (res []int) {
	if root == nil {
		return []int{}
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		size := len(q)
		res = append(res, q[len(q)-1].Val)
		for size > 0 {
			node := q[0]
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
			size--
		}
	}
	return
}
