package problems

/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */

// @lc code=start
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	last, res := 0.0, 1.0
	for res != last {
		last = res
		res = (res + float64(x)/res) / 2
	}
	return int(res)
}

// @lc code=end
