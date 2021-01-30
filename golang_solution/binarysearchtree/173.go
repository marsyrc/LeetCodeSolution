package binarysearchtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	inorder []int
}

func Constructor(root *TreeNode) BSTIterator {
	inorder := []int{}
	s := []*TreeNode{}
	for root != nil || len(s) > 0 {
		for root != nil {
			s = append(s, root)
			root = root.Left
		}
		root = s[len(s)-1]
		s = s[:len(s)-1]
		inorder = append(inorder, root.Val)
		root = root.Right
	}
	return BSTIterator{
		inorder: inorder,
	}
}

func (this *BSTIterator) Next() int {
	next := this.inorder[0]
	this.inorder = this.inorder[1:]
	return next
}

func (this *BSTIterator) HasNext() bool {
	return len(this.inorder) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
