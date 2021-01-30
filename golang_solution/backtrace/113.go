package backtrace

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) (ans [][]int) {
	path := []int{}
	var backTrace func(node *TreeNode, sum int)
	backTrace = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}
		sum -= node.Val
		path = append(path, node.Val)
		defer func() {
			path = path[:len(path)-1]
		}()
		if node.Left == nil && node.Right == nil && sum == 0 {
			ans = append(ans, append([]int(nil), path...))
		}
		backTrace(node.Left, sum)
		backTrace(node.Right, sum)
	}
	backTrace(root, sum)
	return ans
}
