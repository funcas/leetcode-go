package problems

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
	m := map[*ListNode]struct{}{}
	for node := head; node != nil; node = node.Next {
		if _, ok := m[node]; ok {
			return node
		}
		m[node] = struct{}{}
	}

	return nil
}

// @lc code=end
