package problems

/*
 * @lc app=leetcode.cn id=107 lang=golang
 *
 * [107] 二叉树的层次遍历 II
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
type (
	//Queue 队列
	Queue struct {
		top    *node
		rear   *node
		length int
	}
	//双向链表节点
	node struct {
		pre   *node
		next  *node
		value interface{}
	}
)

// Create a new queue
func NewQueue() *Queue {
	return &Queue{nil, nil, 0}
}

//获取队列长度
func (q *Queue) Len() int {
	return q.length
}

//返回true队列不为空
func (q *Queue) Any() bool {
	return q.length > 0
}

//返回队列顶端元素
func (q *Queue) Peek() interface{} {
	if q.top == nil {
		return nil
	}
	return q.top.value
}

//入队操作
func (q *Queue) Push(v interface{}) {
	n := &node{nil, nil, v}
	if q.length == 0 {
		q.top = n
		q.rear = q.top
	} else {
		n.pre = q.rear
		q.rear.next = n
		q.rear = n
	}
	q.length++
}

//出队操作
func (q *Queue) Pop() interface{} {
	if q.length == 0 {
		return nil
	}
	n := q.top
	if q.top.next == nil {
		q.top = nil
	} else {
		q.top = q.top.next
		q.top.pre.next = nil
		q.top.pre = nil
	}
	q.length--
	return n.value
}

func levelOrderBottom(root *TreeNode) [][]int {
	q := NewQueue()
	var ans [][]int
	if root == nil {
		return ans
	}
	q.Push(root)
	for q.Any() {
		var level []int
		size := q.length
		for i := 0; i < size; i++ {
			node := q.Pop().(*TreeNode)
			level = append(level, node.Val)
			if node.Left != nil {
				q.Push(node.Left)
			}
			if node.Right != nil {
				q.Push(node.Right)
			}
		}
		ans = append(ans, level)
	}
	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-1-i] = ans[len(ans)-1-i], ans[i]
	}

	return ans
}

// @lc code=end
