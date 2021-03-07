package binarytree

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

//bfs
func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	queue := []*Node{root}
	for len(queue) > 0 {
		sz := len(queue)
		for i := 0; i < sz; i++ {
			cur := queue[0]
			queue = queue[1:]
			if i < sz-1 {
				cur.Next = queue[0]
			}

			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
	}
	return root
}
