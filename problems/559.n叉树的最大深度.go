package problems

/*
 * @lc app=leetcode.cn id=559 lang=golang
 *
 * [559] N叉树的最大深度
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func maxDepth(root *Node) int {
	ans := 0
	if root == nil {
		return ans
	}
	s := []*Node{root}
	for len(s) > 0 {
		tmp := []*Node{}
		for _, n := range s {
			if len(n.Children) > 0 {
				tmp = append(tmp, n.Children...)
			}
		}
		ans++
		s = tmp
	}
	return ans
}

// @lc code=end
