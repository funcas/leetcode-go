package problems

/*
 * @lc app=leetcode.cn id=258 lang=golang
 *
 * [258] 各位相加
 */
// 计算数根，发现数根 9 个为一组， 1 - 9 循环出现。我们需要做就是把原数映射到树根就可以
// @lc code=start
func addDigits(num int) int {
	return (num-1)%9 + 1
}

// @lc code=end
