package binarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == q.Val || root.Val == p.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */

//并查集方法
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	visited := make(map[int]bool)
	parent := make(map[int]*TreeNode)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			parent[node.Left.Val] = node
			dfs(node.Left)
		}
		if node.Right != nil {
			parent[node.Right.Val] = node
			dfs(node.Right)
		}
	}
	dfs(root)
	for p != nil {
		visited[p.Val] = true
		p = parent[p.Val]
	}

	for q != nil {
		if visited[q.Val] {
			return q
		}
		q = parent[q.Val]
	}
	return nil
}
