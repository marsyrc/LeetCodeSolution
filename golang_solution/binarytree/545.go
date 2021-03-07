package binarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func boundaryOfBinaryTree(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	l := []*TreeNode{}
	if root.Left != nil {
		p := root.Left
		for p.Left != nil || p.Right != nil {
			l = append(l, p)
			if p.Left != nil {
				p = p.Left
				continue
			}
			p = p.Right
		}
	}

	r := []*TreeNode{}
	if root.Right != nil {
		f := root.Right
		for f.Right != nil || f.Left != nil {
			r = append(r, f)
			if f.Right != nil {
				f = f.Right
				continue
			}
			f = f.Left
		}
	}

	leaf := []int{}
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node.Left == nil && node.Right == nil {
			leaf = append(leaf, node.Val)
			return
		}
		if node.Left != nil {
			dfs(node.Left)
		}
		if node.Right != nil {
			dfs(node.Right)
		}
	}
	dfs(root)

	ans := []int{root.Val}
	if root.Left == nil && root.Right == nil {
		return ans
	}
	for _, v := range l {
		ans = append(ans, v.Val)
	}
	ans = append(ans, leaf...)
	for i := len(r) - 1; i >= 0; i-- {
		ans = append(ans, r[i].Val)
	}
	return ans
}
