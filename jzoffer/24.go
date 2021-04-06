package jianzhioffer

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList1(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for head != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return cur
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func reverseList3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head
	for head.Next != nil {
		t := head.Next.Next
		head.Next.Next = cur
		cur = head.Next
		head.Next = t
	}
	return cur
}
