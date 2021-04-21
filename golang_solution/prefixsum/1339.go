package prefixsum

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const mod = 1e9 + 7

var sum []int

func maxProduct(root *TreeNode) int {
	sum = []int{}
	if root == nil {
		return 0
	}
	postOrder(root)
	res := -1
	for i := 0; i < len(sum); i++ {
		res = max(sum[i]*(sum[len(sum)-1]-sum[i]), res)
	}
	return res % mod
}

func postOrder(root *TreeNode) int {
	if root == nil {
		return 0
	}
	cur := postOrder(root.Left) + postOrder(root.Right) + root.Val
	sum = append(sum, cur)
	return cur
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
