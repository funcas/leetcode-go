package problems

/*
 * @lc app=leetcode.cn id=598 lang=golang
 *
 * [598] 范围求和 II
 */

// @lc code=start
func maxCount(m int, n int, ops [][]int) int {
	if len(ops) == 0 {
		return m * n
	}
	minx, miny := 1<<31, 1<<31
	for _, t := range ops {
		if t[0] < minx {
			minx = t[0]
		}
		if t[1] < miny {
			miny = t[1]
		}
	}
	return minx * miny
}

// @lc code=end
