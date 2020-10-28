/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	return preorder(root)
}

func preorder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	rest := append([]int{root.Val}, preorder(root.Left)...)
	rest = append(rest, preorder(root.Right)...)
	return rest
}

// @lc code=end
