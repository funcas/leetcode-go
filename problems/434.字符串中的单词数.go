package problems

import "strings"

/*
 * @lc app=leetcode.cn id=434 lang=golang
 *
 * [434] 字符串中的单词数
 */

// @lc code=start
func CountSegments(s string) int {
	sp := strings.Split(s, " ")
	ans := 0
	for _, v := range sp {
		if v != "" {
			ans++
		}
	}
	return ans
}

// @lc code=end
