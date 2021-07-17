package breadthfirstsearch

type TreeNode struct {
	Val int
    Left *TreeNode
   	Right *TreeNode
}

func widthOfBinaryTree(root *TreeNode) int {
    if root == nil {
        return 0
    }
    q := []*TreeNode{root}
    res := 0

    for len(q) > 0 {
        sz := len(q)
        res = max(res, q[len(q) - 1].Val - q[0].Val + 1)
        for i := 0; i < sz; i++ {
            cur := q[0]
            q = q[1:]
            if cur.Left != nil { 
                cur.Left.Val = cur.Val * 2 + 1
                q = append(q, cur.Left)
            }
            if cur.Right != nil {
                cur.Right.Val = cur.Val * 2 + 2
                q = append(q, cur.Right)
            }
        }
    }
    return res 
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}