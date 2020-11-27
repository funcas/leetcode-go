package problems

import "math"

/*
 * @lc app=leetcode.cn id=147 lang=golang
 *
 * [147] 对链表进行插入排序
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	root := &ListNode{
		Val:  math.MinInt64,
		Next: head,
	}
	prev := root.Next
	cur := head.Next
	for cur != nil {
		tmp := cur
		// tmp.Next = nil
		if tmp.Val < prev.Val {
			prev.Next = cur.Next
			cur = cur.Next
			for p := root; p != nil; p = p.Next {
				if tmp.Val >= p.Val && tmp.Val < p.Next.Val {
					tmp.Next = p.Next
					p.Next = tmp
					break
				}
			}

		} else {
			prev = prev.Next
			cur = cur.Next
		}

	}
	return root.Next

}

// @lc code=end
