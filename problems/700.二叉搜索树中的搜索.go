package problems

/*
 * @lc app=leetcode.cn id=700 lang=golang
 *
 * [700] 二叉搜索树中的搜索
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
func searchBST(root *TreeNode, val int) *TreeNode {
	s := []*TreeNode{root}
	for len(s) > 0 {
		tmp := []*TreeNode{}
		for _, v := range s {
			if v.Val == val {
				return v
			}
			if v.Left != nil {
				tmp = append(tmp, v.Left)
			}
			if v.Right != nil {
				tmp = append(tmp, v.Right)
			}
		}
		s = tmp
	}
	return nil

}

// @lc code=end
