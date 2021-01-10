package binarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	res := []int{}
	var ne *TreeNode

	var dfs func(root, target, father *TreeNode) bool
	dfs = func(root, target, father *TreeNode) bool {
		if root == nil {
			return false
		}
		if root == target {
			ne = father
			return true
		}
		if dfs(root.Left, target, root) {
			root.Left = father
			return true
		}
		if dfs(root.Right, target, root) {
			root.Right = father
			return true
		}
		return false
	}

	var collect func(root *TreeNode, n int, K int)
	collect = func(root *TreeNode, n int, K int) {
		if root == nil {
			return
		}
		if n == K {
			res = append(res, root.Val)
		} else {
			collect(root.Left, n+1, K)
			collect(root.Right, n+1, K)
		}
		return
	}

	dfs(root, target, nil)
	collect(target, 0, K)
	collect(ne, 0, K-1)
	return res
}
