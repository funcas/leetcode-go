package problems

/*
 * @lc app=leetcode.cn id=371 lang=golang
 *
 * [371] 两整数之和
 */
// 1 a ^ b 计算没有进位的和
// 2 (a & b)<<1 计算进位
// 3 进位与没进位的和相加，判断是否还有进位，有就重复，没有就退出。

// GetSum 正整数二进制为原码，负数二进制为补码， 即反码+1 -3 => ^-3 + 1。
// @lc code=start
func getSum(a int, b int) int {
	for b != 0 {
		t := a ^ b
		b = (a & b) << 1
		a = t
	}

	return a
}

// @lc code=end
