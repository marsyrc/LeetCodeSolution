package jianzhioffer

func verifyPostorder(postorder []int) bool {
	if len(postorder) == 0 {
		return true
	}
	i := 0
	for postorder[i] < postorder[len(postorder)-1] {
		i++
	}
	mark := i
	for postorder[i] > postorder[len(postorder)-1] {
		i++
	}
	if i != len(postorder)-1 {
		return false
	}
	return verifyPostorder(postorder[:mark]) && verifyPostorder(postorder[mark:len(postorder)-1])
}
