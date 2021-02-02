/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第N个节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// func removeNthFromEnd(head *ListNode, n int) *ListNode {
// 	dummy := &ListNode{}
// 	dummy.Next = head
// 	p := head
// 	length := 0
// 	for p != nil {
// 		length++
// 		p = p.Next
// 	}
// 	length -= n
// 	p = dummy
// 	for length > 0 {
// 		length--
// 		p = p.Next
// 	}
// 	p.Next = p.Next.Next
// 	return dummy.Next
// }

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	p := dummy
	q := dummy
	for i := 0; i <= n; i++ {
		p = p.Next
	}

	for p != nil {
		p = p.Next
		q = q.Next
	}
	q.Next = q.Next.Next
	return dummy.Next
}

// @lc code=end
