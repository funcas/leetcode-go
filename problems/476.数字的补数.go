package problems

import "fmt"

/*
 * @lc app=leetcode.cn id=476 lang=golang
 *
 * [476] 数字的补数
 */

// @lc code=start
func findComplement(num int) int {
	a := fmt.Sprintf("%b", num)
	b := []byte(a)
	for i := 0; i < len(b); i++ {
		b[i] = b[i] ^ 1
	}
	s := string(b)
	l := len(s)
	ans := 0
	for i := l - 1; i >= 0; i-- {
		ans += (int(s[l-i-1]) & 0xf) << uint8(i)
	}
	return ans
}

// @lc code=end
