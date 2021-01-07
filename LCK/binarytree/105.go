package binarysearchtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//too slow
func buildTree(preorder []int, inorder []int) *TreeNode {
	place := make(map[int]int)
	for i, v := range inorder {
		place[v] = i
	}

	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{preorder[0], nil, nil}
	i := place[preorder[0]]
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}
