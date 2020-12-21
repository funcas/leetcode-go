package problems

/*
 * @lc app=leetcode.cn id=653 lang=golang
 *
 * [653] 两数之和 IV - 输入 BST
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
	var mid func(root *TreeNode)
	arr := []int{}
	mid = func(root *TreeNode) {
		if root == nil {
			return
		}
		mid(root.Left)
		arr = append(arr, root.Val)
		mid(root.Right)
	}
	mid(root)
	for i, j := 0, len(arr)-1; i < j; {
		if arr[i]+arr[j] == k {
			return true
		}
		if arr[i]+arr[j] < k {
			i++
		} else {
			j--
		}
	}
	return false
}

// @lc code=end
