/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {

	nodeMap := make(map[*ListNode]int)
	for head != nil {
		_, ok := nodeMap[head]
		if ok {
			return head
		} else {
			nodeMap[head] = 0
		}
		head = head.Next
	}
	return nil

}

// @lc code=end
