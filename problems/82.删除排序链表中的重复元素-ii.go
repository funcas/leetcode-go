package problems

/*
 * @lc app=leetcode.cn id=82 lang=golang
 *
 * [82] 删除排序链表中的重复元素 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	root := &ListNode{
		Next: head,
	}

	prev, cur := root, root.Next
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			for cur.Next != nil && cur.Next.Val == cur.Val {
				cur = cur.Next
			}
			prev.Next = cur.Next
			cur.Next = nil
			cur = prev.Next
		} else {
			cur = cur.Next
			prev = prev.Next
		}
	}
	return root.Next
}

// @lc code=end
