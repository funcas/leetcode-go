package problems

/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
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
func maxDepth104(root *TreeNode) int {
	maxDepth := 0
	if root == nil {
		return maxDepth
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		tmp := make([]*TreeNode, 0)
		for _, v := range queue {

			if v.Left != nil {
				tmp = append(tmp, v.Left)
			}
			if v.Right != nil {
				tmp = append(tmp, v.Right)
			}
		}
		queue = tmp
		maxDepth++
	}
	return maxDepth
}

// @lc code=end
