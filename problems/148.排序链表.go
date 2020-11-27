package problems

/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func sortList(head *ListNode) *ListNode {
	return partation1(head, nil)
}

func partation1(head *ListNode, tail *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == tail {
		head.Next = nil
		return head
	}
	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}
	mid := slow
	left := partation1(head, mid)
	right := partation1(mid, tail)
	return merge1(left, right)

}

func merge1(left *ListNode, right *ListNode) *ListNode {
	root := &ListNode{}
	m, p, q := root, left, right
	for p != nil && q != nil {
		if p.Val <= q.Val {
			m.Next = p
			p = p.Next
		} else {
			m.Next = q
			q = q.Next
		}
		m = m.Next
	}
	if p != nil {
		m.Next = p
	} else if q != nil {
		m.Next = q
	}
	return root.Next
}

// @lc code=end
