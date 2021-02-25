package backtrace

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pseudoPalindromicPaths(root *TreeNode) (res int) {
	oddSet := make(map[int]int)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		_, ok := oddSet[node.Val]
		if ok {
			delete(oddSet, node.Val)
			if node.Left == nil && node.Right == nil {
				if len(oddSet) <= 1 {
					res++
				}
			}
			dfs(node.Left)
			dfs(node.Right)
			oddSet[node.Val]++
		} else {
			oddSet[node.Val]++
			if node.Left == nil && node.Right == nil {
				if len(oddSet) <= 1 {
					res++
				}
			}
			dfs(node.Left)
			dfs(node.Right)
			delete(oddSet, node.Val)
		}
	}
	dfs(root)
	return res
}
