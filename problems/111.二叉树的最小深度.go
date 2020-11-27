package problems

import "math"

/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
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
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	minDep := math.MaxInt32
	if root.Left != nil {
		minDep = min(minDepth(root.Left), minDep)
	}
	if root.Right != nil {
		minDep = min(minDepth(root.Right), minDep)
	}

	return minDep + 1
}

// @lc code=end
