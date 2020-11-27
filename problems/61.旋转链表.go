package problems

/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	count := 0
	for p := head; p != nil; p = p.Next {
		count++
	}
	if count == 0 {
		return head
	}
	times := k % count

	m, n := head, head
	for i := 0; i < times; i++ {
		m = m.Next
	}

	for m.Next != nil {
		m = m.Next
		n = n.Next
	}

	m.Next = head
	cur := n.Next
	n.Next = nil

	return cur
}

// @lc code=end
