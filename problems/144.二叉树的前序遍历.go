package problems

import "github.com/funcas/algorithm/ds"

/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
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
func preorderTraversal(root *TreeNode) []int {
	ans := []int{}
	var preorder func(node *TreeNode)

	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		ans = append(ans, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return ans
}

func preorderTraversal2(root *TreeNode) []int {
	S := ds.NewStack()
	node := root
	var ans []int
	for node != nil || !S.Empty() {
		for node != nil {
			ans = append(ans, node.Val)
			S.Push(node)
			node = node.Left
		}
		if !S.Empty() {
			tmp := S.Pop().(*TreeNode)
			node = tmp.Right
		}
	}
	return ans
}

// @lc code=end
