package problems

import (
	"strings"
)

/*
 * @lc app=leetcode.cn id=290 lang=golang
 *
 * [290] 单词规律
 */

// @lc code=start
func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(words) != len(pattern) {
		return false

	}
	m := make(map[string]uint8)
	m2 := make(map[uint8]string)
	for i := 0; i < len(words); i++ {
		if v, ok := m[words[i]]; ok {
			if v != pattern[i] {
				return false
			}
		} else {
			m[words[i]] = pattern[i]
		}

	}
	for i := 0; i < len(pattern); i++ {
		if v, ok := m2[pattern[i]]; ok {
			if v != words[i] {
				return false
			}
		} else {
			m2[pattern[i]] = words[i]
		}
	}
	return true
}

// @lc code=end
