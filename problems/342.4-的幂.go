package problems

/*
 * @lc app=leetcode.cn id=342 lang=golang
 *
 * [342] 4的幂
 */

// @lc code=start
func isPowerOfFour(num int) bool {
	if num == 0 {
		return false
	}
	for num%4 == 0 {
		num = num / 4
	}
	return num == 1
}

// @lc code=end
