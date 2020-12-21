package problems

import (
	"funcas.github.com/leetcode/uncommit"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=606 lang=golang
 *
 * [606] 根据二叉树创建字符串
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
func Tree2str(t *uncommit.TreeNode) string {

	if t == nil {
		return ""
	}
	if t.Left == nil && t.Right == nil {
		return strconv.Itoa(t.Val)
	}
	if t.Right == nil {
		return strconv.Itoa(t.Val) + "(" + Tree2str(t.Left) + ")"
	}

	return strconv.Itoa(t.Val) + "(" + Tree2str(t.Left) + ")" + "(" + Tree2str(t.Right) + ")"
}

// @lc code=end
