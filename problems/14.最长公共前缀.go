package problems

import "strings"

/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	p := strs[0]
	for i := 1; i < len(strs); i++ {
		for strings.Index(strs[i], p) != 0 {
			p = p[0 : len(p)-1]
			if len(p) == 0 {
				return ""
			}
		}
	}
	return p
}

// @lc code=end
