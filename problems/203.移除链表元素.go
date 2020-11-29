package problems

/*
 * @lc app=leetcode.cn id=203 lang=golang
 *
 * [203] 移除链表元素
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	// 加一个空头，这样处理第一个元素跟处理其它元素就一样了
	tmp := &ListNode{
		Next: head,
	}
	// 当前指针
	c := tmp

	for c.Next != nil {
		if c.Next.Val == val {
			c.Next = c.Next.Next
			continue
		}
		c = c.Next
	}
	return tmp.Next
}

// @lc code=end
