package problems

/*
 * @lc app=leetcode.cn id=108 lang=golang
 *
 * [108] 将有序数组转换为二叉搜索树
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
func sortedArrayToBST(nums []int) *TreeNode {
	size := len(nums)
	if size == 0 {
		return nil
	}
	var buildBST func(l, r int) *TreeNode
	buildBST = func(l, r int) *TreeNode {
		if l > r {
			return nil
		}

		m := (l + r) / 2
		root := &TreeNode{Val: nums[m]}
		root.Left = buildBST(l, m-1)
		root.Right = buildBST(m+1, r)

		return root
	}
	return buildBST(0, size-1)

}

// @lc code=end
