package binarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//preorder ergodic with stack
func flatten(root *TreeNode) {
	list := []*TreeNode{}
	stack := []*TreeNode{}

	p := root
	for p != nil || len(stack) > 0 {
		for p != nil {
			list = append(list, p)
			stack = append(stack, p)
			p = p.Left
		}

		p = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		p = p.Right
	}

	for i := 1; i < len(list); i++ {
		pre, cur := list[i-1], list[i]
		pre.Left, pre.Right = nil, cur
	}
}
