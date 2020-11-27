package problems

/*
 * @lc app=leetcode.cn id=461 lang=golang
 *
 * [461] 汉明距离
 */

// @lc code=start
func hammingDistance(x int, y int) int {
	r := x ^ y
	count := 0
	for r != 0 {
		r = r & (r - 1)
		count++
	}

	return count

}

// @lc code=end
