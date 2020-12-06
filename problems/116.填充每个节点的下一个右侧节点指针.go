package problems

/*
 * @lc app=leetcode.cn id=116 lang=golang
 *
 * [116] 填充每个节点的下一个右侧节点指针
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node2) *Node2 {
	if root == nil {
		return nil
	}
	queue := []*Node2{root}

	for len(queue) != 0 {
		size := len(queue)
		tmp := make([]*Node2, 0)
		for i := 0; i < size; i++ {
			if i+1 < size {
				queue[i].Next = queue[i+1]
			}

			if queue[i].Left != nil {
				tmp = append(tmp, queue[i].Left)
			}
			if queue[i].Right != nil {
				tmp = append(tmp, queue[i].Right)
			}
		}

		queue = tmp

	}
	return root
}

// @lc code=end
