package listnode

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	less, great := &ListNode{}, &ListNode{}
	l, g := less, great
	for head != nil {
		if head.Val < x {
			l.Next = head
			l = l.Next
		} else {
			g.Next = head
			g = g.Next
		}
		head = head.Next
	}
	//connect
	l.Next = great.Next
	g.Next = nil
	return less.Next
}
