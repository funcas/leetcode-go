package problems

/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] 二叉树展开为链表
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
func flatten(root *TreeNode) {

	var cur *TreeNode
	var lastOrder func(node *TreeNode)

	lastOrder = func(node *TreeNode) {
		if node == nil {
			return
		}

		lastOrder(node.Right)
		lastOrder(node.Left)
		node.Right = cur
		node.Left = nil
		cur = node
	}
	lastOrder(root)

}

// @lc code=end
