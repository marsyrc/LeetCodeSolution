package listnode

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{}
	dummy.Next = head
	slow := dummy
	fast := head
	for fast != nil && fast.Next != nil {
		if slow.Next.Val != fast.Next.Val {
			slow = slow.Next
			fast = fast.Next
		} else {
			for fast != nil && fast.Next != nil && slow.Next.Val == fast.Next.Val {
				fast = fast.Next
			}
			slow.Next = fast.Next
			fast = fast.Next
		}
	}
	return dummy.Next
}
