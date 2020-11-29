package problems

/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	root := &ListNode{
		Val:  0,
		Next: head,
	}

	for p := root.Next; p != nil && p.Next != nil; {
		tmp := p.Next
		p.Next = p.Next.Next
		tmp.Next = root.Next
		root.Next = tmp
	}
	return root.Next
}

// @lc code=end
