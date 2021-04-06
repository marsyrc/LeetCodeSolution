package jianzhioffer

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func treeToDoublyList(root *Node) *Node {
	var pre *Node
	var head *Node
	var dfs func(cur *Node)
	dfs = func(cur *Node) {
		if cur == nil {
			return
		}
		dfs(cur.Left)
		if pre != nil {
			pre.Right = cur
		} else {
			head = cur
		}
		cur.Left = pre
		pre = cur
		dfs(cur.Right)
	}
	if root == nil {
		return nil
	}
	dfs(root)
	head.Left = pre
	pre.Right = head
	return head
}

func treeToDoublyList2(root *Node) *Node {
	list := []*Node{}
	if root == nil {
		return nil
	}
	var inorder func(node *Node)
	inorder = func(node *Node) {
		if node == nil {
			return
		}
		inorder(node.Left)
		list = append(list, node)
		inorder(node.Right)
	}
	inorder(root)
	head := &Node{0, nil, list[0]}
	list[0].Left = list[len(list)-1]
	list[len(list)-1].Right = list[0]
	for i := range list {
		if i < len(list)-1 {
			list[i].Right = list[i+1]
		}
		if i > 0 {
			list[i].Left = list[i-1]
		}
	}
	return head.Right
}
