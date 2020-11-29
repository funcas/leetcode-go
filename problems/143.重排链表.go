package problems

/*
 * @lc app=leetcode.cn id=143 lang=golang
 *
 * [143] 重排链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	tmp := []*ListNode{}
	for p := head; p != nil; p = p.Next {
		tmp = append(tmp, p)
	}

	i, j := 0, len(tmp)-1
	for i < j {
		tmp[i].Next = tmp[j]
		i++
		if i == j {
			break
		}
		tmp[j].Next = tmp[i]
		j--
	}
	tmp[i].Next = nil
}

// @lc code=end
