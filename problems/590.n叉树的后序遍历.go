package problems

/*
 * @lc app=leetcode.cn id=590 lang=golang
 *
 * [590] N叉树的后序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
	ans := []int{}

	var post func(node *Node)
	post = func(node *Node) {
		if node == nil {
			return
		}
		for _, n := range node.Children {
			post(n)
		}
		ans = append(ans, node.Val)
	}
	post(root)
	return ans
}

// @lc code=end
