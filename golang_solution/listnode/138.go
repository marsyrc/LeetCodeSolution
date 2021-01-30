package listnode

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	visited := map[*Node]*Node{}
	getCloned := func(old *Node) *Node {
		if n, ok := visited[old]; ok {
			return n
		} else {
			if old == nil {
				return nil
			}
			newNode := &Node{old.Val, nil, nil}
			visited[old] = newNode
			return newNode
		}
	}
	new := getCloned(head)
	newHead := new
	for head != nil {
		new.Next = getCloned(head.Next)
		new.Random = getCloned(head.Random)
		new = new.Next
		head = head.Next
	}
	return newHead
}
