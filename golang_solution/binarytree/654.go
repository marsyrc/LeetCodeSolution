package binarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//recursion solution
func constructMaximumBinaryTree(nums []int) *TreeNode {
	var f func(start int, end int) *TreeNode
	f = func(start int, end int) *TreeNode {
		if start > end {
			return nil
		}
		i := findMaxsIndex(nums, start, end)
		node := &TreeNode{
			Val:   nums[i],
			Left:  f(start, i-1),
			Right: f(i+1, end),
		}
		return node
	}
	return f(0, len(nums)-1)
}

func findMaxsIndex(nums []int, l int, r int) int {
	res := l
	for i := l + 1; i <= r; i++ {
		if nums[i] > nums[res] {
			res = i
		}
	}
	return res
}
