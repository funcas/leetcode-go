package problems

/*
 * @lc app=leetcode.cn id=172 lang=golang
 *
 * [172] 阶乘后的零
 */

// @lc code=start
func TrailingZeroes(n int) int {
	ans := 0
	for n >= 5 {
		ans += n / 5
		n /= 5
	}
	return ans
}

// @lc code=end
