package problems

/*
 * @lc app=leetcode.cn id=441 lang=golang
 *
 * [441] 排列硬币
 */

// @lc code=start
func arrangeCoins(n int) int {
	if n == 0 {
		return 0
	}
	level := 2
	for (level*(level+1))/2 <= n {
		level++
	}
	return level - 1
}

// @lc code=end
