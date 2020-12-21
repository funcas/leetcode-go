package problems

/*
 * @lc app=leetcode.cn id=589 lang=golang
 *
 * [589] N叉树的前序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
	ans := []int{}
	var pre func(node *Node)
	pre = func(node *Node) {
		if node == nil {
			return
		}
		ans = append(ans, node.Val)
		for _, n := range node.Children {
			pre(n)
		}
	}
	pre(root)
	return ans
}

// @lc code=end
