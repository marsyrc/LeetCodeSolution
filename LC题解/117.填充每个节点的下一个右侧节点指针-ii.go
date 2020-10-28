import "fmt"

/*
 * @lc app=leetcode.cn id=117 lang=golang
 *
 * [117] 填充每个节点的下一个右侧节点指针 II
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	q := []*Node{root}
	for len(q) > 0 {
		sz := len(q)
		for i := 0; i < sz; i++ {

			node := q[0]
			q = q[1:] //queue.poll

			if i < sz-1 {
				node.Next = q[0]
			} else {
				node.Next = nil
			}

			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		fmt.Print("q length is", len(q))
	}
	return root
}

// @lc code=end
