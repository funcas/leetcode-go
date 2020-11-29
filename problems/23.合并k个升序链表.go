package problems

/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个升序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	l := len(lists)
	if l == 0 {
		return nil
	}
	if l == 1 {
		return lists[0]
	}
	ret := mergeList(lists[0], lists[1])
	for i := 2; i < len(lists); i++ {
		ret = mergeList(ret, lists[i])
	}
	return ret
}

func mergeList(a *ListNode, b *ListNode) *ListNode {
	c := &ListNode{}
	tail := c
	p := a
	q := b
	for p != nil && q != nil {
		if p.Val < q.Val {
			tail.Next = p
			p = p.Next
		} else {
			tail.Next = q
			q = q.Next
		}
		tail = tail.Next
	}
	if p != nil {
		tail.Next = p
	}
	if q != nil {
		tail.Next = q
	}
	return c.Next
}

// @lc code=end
