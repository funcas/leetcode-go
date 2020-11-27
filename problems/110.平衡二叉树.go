package problems

/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
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
func isBalanced(root *TreeNode) bool {
	var balance func(node *TreeNode) int
	balance = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := balance(node.Left)
		if left == -1 {
			return -1
		}
		right := balance(node.Right)
		if right == -1 {
			return -1
		}
		if abs(left-right) <= 1 {
			return max(left, right) + 1
		}
		return -1

	}
	return balance(root) != -1
}

// @lc code=end
