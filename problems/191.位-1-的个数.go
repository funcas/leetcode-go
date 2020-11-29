package problems

import "fmt"

/*
 * @lc app=leetcode.cn id=191 lang=golang
 *
 * [191] 位1的个数
 */

// @lc code=start
func hammingWeight(num uint32) int {
	c := 0
	a := fmt.Sprintf("%b", num)
	for i := 0; i < len(a); i++ {
		if a[i] == 49 {
			c++
		}
	}
	return c
}

func hammingWeight2(num uint32) int {
	c := 0
	var a uint32 = 1
	for i := 0; i < 32; i++ {
		if (num & a) != 0 {
			c++
		}
		a = a << 1
	}
	return c
}

// @lc code=end
