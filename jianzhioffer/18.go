package jianzhioffer

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{0, head}
	pre, cur := dummyHead, head
	for cur.Val != val {
		pre = pre.Next
		cur = cur.Next
	}
	pre.Next = cur.Next
	return dummyHead.Next
}
