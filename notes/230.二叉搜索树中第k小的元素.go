/*
 * @lc app=leetcode.cn id=230 lang=golang
 *
 * [230] 二叉搜索树中第K小的元素
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
//递归写法，中序遍历
func kthSmallest(root *TreeNode, k int) int {
	var ansArray []int

	inorder(root, &ansArray)

	return ansArray[k-1]
}
func inorder(root *TreeNode, ansArray *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, ansArray)
	*ansArray = append(*ansArray, root.Val)
	inorder(root.Right, ansArray)
	return
}

//利用栈进行迭代，中序遍历
/*
func kthSmallest(root *TreeNode, k int) int {
	var stack []*TreeNode
	for true{
		for root != nil{
			stack = append(stack,root)
			root = root.Left
		}
		if len(stack) == 0{
			return 0
		}
		root = stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		k--
		if k == 0{
			return root.Val
		}
		root = root.Right
	}
	return 0
}
*/

//二分
/*
func kthSmallest(root *TreeNode, k int) int {
	var count int

	count = sumRoot(root.Left)

	if k < 1+ count {
		return kthSmallest(root.Left, k)
	}else if k > count + 1{
		return kthSmallest(root.Right,k-count - 1)
	}
	return root.Val
}

func sumRoot(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + sumRoot(root.Left) + sumRoot(root.Left)
}

*/

// @lc code=end
