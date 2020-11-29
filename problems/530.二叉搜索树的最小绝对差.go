package problems

/*
 * @lc app=leetcode.cn id=530 lang=golang
 *
 * [530] 二叉搜索树的最小绝对差
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
func getMinimumDifference(root *TreeNode) int {
	ans := 1<<31 - 1
	pre := -1
	midOrder(root, &pre, &ans)
	return ans
}

func midOrder(node *TreeNode, tmp *int, res *int) {

	if node == nil {
		return
	}

	midOrder(node.Left, tmp, res)
	min := node.Val - *tmp
	if *tmp != -1 && min-*res < 0 {
		*res = min
	}
	*tmp = node.Val
	midOrder(node.Right, tmp, res)
}

// @lc code=end
