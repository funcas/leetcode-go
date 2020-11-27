package problems

import "math"

/*
 * @lc app=leetcode.cn id=414 lang=golang
 *
 * [414] 第三大的数
 */

// @lc code=start
func thirdMax(nums []int) int {
	const inf = -math.MaxFloat64
	a, b, c := inf, inf, inf
	for _, v := range nums {
		tmp := float64(v)
		if tmp > a {
			c = b
			b = a
			a = tmp
		} else if tmp == a {
			continue
		} else if tmp > b {
			c = b
			b = tmp
		} else if tmp == b {
			continue
		} else if tmp > c {
			c = tmp
		}
	}
	if c != inf {
		return int(c)
	}
	return int(a)
}

// @lc code=end
