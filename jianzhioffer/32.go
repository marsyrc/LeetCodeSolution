package jianzhioffer

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 32-I
func levelOrder(root *TreeNode) (res []int) {
	if root == nil {
		return res
	}
	// 根节点入队
	q := []*TreeNode{root}
	var node *TreeNode
	for i := 0; i < len(q); i++ {
		node = q[i]
		res = append(res, node.Val)
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
	}
	return res
}

//32-II
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder2(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		sz := len(q)
		level := []int{}
		for i := 0; i < sz; i++ {
			node := q[0]
			q = q[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append(res, level)
	}
	return res
}

//32-III
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder3(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var ans [][]int
	queue := []*TreeNode{root}
	fromLToR := true
	for len(queue) != 0 {
		sz := len(queue)
		level := []int{}
		for i := 0; i < sz; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if fromLToR {
			ans = append(ans, level)
		} else {
			reverse(level)
			ans = append(ans, level)
		}
		fromLToR = !fromLToR
	}
	return ans
}

func reverse(a []int) {
	if len(a) <= 1 {
		return
	}
	l, r := 0, len(a)-1
	for l < r {
		a[l], a[r] = a[r], a[l]
		l++
		r--
	}
	return
}
