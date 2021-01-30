package binarysearchtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return getTree(0, n-1)
}

func getTree(start int, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	curTrees := []*TreeNode{}
	for i := start; i <= end; i++ {
		leftTrees := getTree(start, i-1)
		rightTrees := getTree(i+1, end)
		for _, l := range leftTrees {
			for _, r := range rightTrees {
				curNode := &TreeNode{
					Val:   i + 1,
					Left:  l,
					Right: r,
				}
				curTrees = append(curTrees, curNode)
			}
		}
	}
	return curTrees
}
