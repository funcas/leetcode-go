package problems

import "sort"

/*
 * @lc app=leetcode.cn id=720 lang=golang
 *
 * [720] 词典中最长的单词
 */

// @lc code=start
func LongestWord(words []string) string {
	sort.Strings(words)
	m := make(map[string]struct{})
	ans := ""
	for _, w := range words {
		if _, ok := m[w[:len(w)-1]]; len(w) == 1 || ok {
			if len(ans) < len(w) {
				ans = w
			}
			m[w] = struct{}{}
		}
	}
	return ans
}

// @lc code=end
