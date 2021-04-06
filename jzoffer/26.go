package jianzhioffer

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	var recursion func(A *TreeNode, B *TreeNode) bool
	recursion = func(A *TreeNode, B *TreeNode) bool {
		if B == nil {
			return true
		}
		if A == nil || A.Val != B.Val {
			return false
		}
		return recursion(A.Left, B.Left) && recursion(A.Right, B.Right)
	}
	return A != nil && B != nil && (recursion(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B))
}
