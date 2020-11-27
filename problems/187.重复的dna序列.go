package problems

import "fmt"

/*
 * @lc app=leetcode.cn id=187 lang=golang
 *
 * [187] 重复的DNA序列
 */

// @lc code=start
func findRepeatedDnaSequences(s string) []string {
	m := make(map[string]int)
	for i, j := 0, 10; j <= len(s); i, j = i+1, j+1 {
		ss := fmt.Sprintf("%s", s[i:j])
		if _, ok := m[ss]; ok {
			m[ss]++
		} else {
			m[ss] = 1
		}
	}
	ans := []string{}
	for k, v := range m {
		if v > 1 {
			ans = append(ans, k)
		}
	}
	return ans
}

// @lc code=end
