package problems

/*
 * @lc app=leetcode.cn id=693 lang=golang
 *
 * [693] 交替位二进制数
 */

// @lc code=start
func hasAlternatingBits(n int) bool {
	for n > 0 {
		if n%2 == (n/2)%2 {
			return false
		}
		n /= 2
	}
	return true
}

// 右移一位异或，如果间隔都不相同，那么势必异或后的值全1
// 那么只要判断这个结果是否全1即可
func hasAlternatingBits2(n int) bool {
	tmp := n ^ (n >> 1)
	return tmp&(tmp+1) == 0
}

// @lc code=end
