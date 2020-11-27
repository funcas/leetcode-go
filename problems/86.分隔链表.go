package problems

/*
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] 分隔链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	root := &ListNode{
		Next: head,
	}
	p := root
	q := root
	for p.Next != nil && p.Next.Val < x {
		p = p.Next
		q = p
	}

	for p.Next != nil {
		if p.Next.Val < x {
			tmp := p.Next
			p.Next = p.Next.Next
			tmp.Next = q.Next
			q.Next = tmp
			q = q.Next
		} else {
			p = p.Next
		}
	}
	return root.Next
}

// @lc code=end
