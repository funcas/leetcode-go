package problems

/*
 * @lc app=leetcode.cn id=328 lang=golang
 *
 * [328] 奇偶链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	i := 2
	p := head.Next
	cur := head
	prev := head
	for p != nil {
		tmp := p.Next
		if i%2 != 0 {
			prev.Next = tmp
			p.Next = cur.Next
			cur.Next = p
			cur = cur.Next
		} else {
			prev = p
		}

		p = tmp
		i++
	}
	return head
}

// @lc code=end
