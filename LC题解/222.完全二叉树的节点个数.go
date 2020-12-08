/*
 * @lc app=leetcode.cn id=222 lang=golang
 *
 * [222] 完全二叉树的节点个数
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
/**
对任意一个子树，遍历其左子树层高left，
右子树层高right，相等左子树则是满二叉树，
否则右子树是满二叉树。
*/
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := countLevel(root.Left)
	right := countLevel(root.Right)
	if right == left {
		return countNodes(root.Right) + (1 << left)
	} else {
		return countNodes(root.Left) + (1 << right)
	}
}

func countLevel(root *TreeNode) int {
	level := 0
	for root != nil {
		level++
		root = root.Left
	}
	return level
}

// @lc code=end
