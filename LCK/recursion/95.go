package recursion

//TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	var res []*TreeNode

	var getTree func(start int, end int) []*TreeNode
	getTree = func(start int, end int) (ans []*TreeNode) {
		if start > end {
			return []*TreeNode{nil}
		}
		if start == end {
			ans = append(ans, &TreeNode{Val: end})
			return
		}
		for j := start; j <= end; j++ {
			lefts := getTree(start, j-1)
			rights := getTree(j+1, end)
			for _, v := range lefts {
				for _, u := range rights {
					root := &TreeNode{
						Val:   j,
						Left:  v,
						Right: u,
					}
					ans = append(ans, root)
				}
			}
		}
		return
	}
	res = getTree(1, n)
	return res
}
