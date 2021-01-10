package binarysearchtree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	s := []*TreeNode{}
	var x, y, pre *TreeNode
	for len(s) > 0 || root != nil {
		for root != nil {
			s = append(s, root)
			root = root.Left
		}
		root = s[len(s)-1]
		s = s[:len(s)-1]
		if x == nil && pre != nil && root.Val < pre.Val {
			x = pre
		}
		if x != nil && pre != nil && root.Val < pre.Val {
			y = root
			fmt.Println(y.Val)
		}
		pre = root
		root = root.Right
	}
	x.Val, y.Val = y.Val, x.Val
	return
}
