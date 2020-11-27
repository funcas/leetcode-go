package problems

/*
 * @lc app=leetcode.cn id=222 lang=golang
 *
 * [222] 完全二叉树的节点个数
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
func countNodes(root *TreeNode) int {
	count := 0
	var order func(node *TreeNode)
	order = func(node *TreeNode) {
		if node == nil {
			return
		}

		count++
		order(node.Left)
		order(node.Right)
	}

	order(root)
	return count

}

// @lc code=end
