package problems

/*
 * @lc app=leetcode.cn id=141 lang=golang
 *
 * [141] 环形链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	m := map[*ListNode]int{}
	for node := head; node != nil; node = node.Next {
		if _, ok := m[node]; ok {
			return true
		}
		m[node] = node.Val
	}

	return false
}

// @lc code=end
