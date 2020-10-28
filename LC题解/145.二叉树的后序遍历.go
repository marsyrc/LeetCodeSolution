/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
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
func postorderTraversal(root *TreeNode) []int {

	//存储结果
	var res []int
	if root == nil {
		return res
	}
	//通过栈控制遍历顺序
	stack := []*TreeNode{root}
	//用map标记某个节点是否 遍历过
	marks := make(map[*TreeNode]bool)

	var node *TreeNode
	for len(stack) != 0 {
		//pop最后一个节点
		i := len(stack) - 1
		node = stack[i]
		stack = stack[:i]

		//如果节点已经遍历过，则打印
		_, ok := marks[node]
		if ok {
			res = append(res, node.Val)
		} else {
			//否则，入栈该节点和右孩子，左孩子，并标记当前节点已经遍历过(下次不再处理，可直接打印)

			//TODO 更改下行代码的位置 可以转换 前序，中序，后序遍历(当前为后序遍历)
			stack = append(stack, node)

			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			//TODO 上述代码放到这里即 中序遍历
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			//TODO 上述代码放到这里即 前序遍历
			marks[node] = true
		}
	}
	return res
}

// @lc code=end
