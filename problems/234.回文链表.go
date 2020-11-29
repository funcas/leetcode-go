package problems

/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome_(head *ListNode) bool {
	tmp := []int{}
	for p := head; p != nil; p = p.Next {
		tmp = append(tmp, p.Val)
	}
	for i, j := 0, len(tmp)-1; i < j; i, j = i+1, j-1 {
		if tmp[i] != tmp[j] {
			return false
		}
	}
	return true
}
func isPalindrome2_(head *ListNode) bool {
	var s1 int64
	var s2 int64
	t := 1
	s1, s2 = 0, 0
	for head != nil {
		s1 = s1*10 + int64(head.Val)
		s2 = s2 + int64(t*head.Val)
		t = t * 10
		head = head.Next
	}
	return s1 == s2
}

// @lc code=end
