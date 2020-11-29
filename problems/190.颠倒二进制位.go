package problems

import "fmt"

/*
 * @lc app=leetcode.cn id=190 lang=golang
 *
 * [190] 颠倒二进制位
 */

// @lc code=start
func reverseBits(num uint32) uint32 {
	x := fmt.Sprintf("%032b", num)
	var r uint32 = 0
	for i := 0; i < len(x); i++ {
		r += (uint32(x[i]) & 0xf) << uint8(i)
	}
	return r
}

// @lc code=end
