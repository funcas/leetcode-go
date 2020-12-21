package problems

/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func ReverseBetween(head *ListNode, m int, n int) *ListNode {
	root := &ListNode{
		Next: head,
	}
	i := 0
	pr, p, q := root, root, root
	for i <= n && p != nil {
		if i+1 == m {
			pr = p
		}

		if i > m {
			q.Next = p.Next
			p.Next = pr.Next
			pr.Next = p
			p = q.Next
		} else {
			q = p
			p = p.Next
		}

		i++
	}
	return root.Next
}

// @lc code=end
/*
  pr q  p
  1  2  3  4  5
  1  3  2  4  5
  pr p  q
  1  4  3  2  5
  pr q  p


*/
