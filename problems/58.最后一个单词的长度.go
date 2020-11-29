package problems

import "strings"

/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	s = strings.Trim(s, " ")
	str := strings.Split(s, " ")
	size := len(str)
	if size == 0 {
		return 0
	}
	return len(str[len(str)-1])
}

// @lc code=end
