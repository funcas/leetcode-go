package problems

import "strconv"

/*
 * @lc app=leetcode.cn id=1323 lang=golang
 *
 * [1323] 6 和 9 组成的最大数字
 */

// @lc code=start
func Maximum69Number(num int) int {
	ns := []byte(strconv.Itoa(num))
	for i := range ns {
		if ns[i] == '6' {
			ns[i] = '9'
			break
		}
	}
	ans, _ := strconv.Atoi(string(ns))
	return ans
}

// @lc code=end
