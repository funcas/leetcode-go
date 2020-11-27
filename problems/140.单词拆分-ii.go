package problems

import "strings"

/*
 * @lc app=leetcode.cn id=140 lang=golang
 *
 * [140] 单词拆分 II
 */

// @lc code=start
func wordBreak(s string, wordDict []string) []string {
	m := make(map[string]struct{})
	for _, v := range wordDict {
		m[v] = struct{}{}
	}
	var dfs func(start int) [][]string
	size := len(s)
	memo := make([][][]string, size)
	dfs = func(start int) [][]string {
		if memo[start] != nil {
			return memo[start]
		}
		res := [][]string{}
		for i := start + 1; i < size; i++ {
			word := s[start:i]
			if _, ok := m[word]; ok {
				retainWord := dfs(i)
				for _, rw := range retainWord {
					res = append(res, append([]string{word}, rw...))
				}
			}

		}
		word := s[start:]
		if _, ok := m[word]; ok {
			res = append(res, []string{word})
		}
		memo[start] = res
		return res
	}

	ans := []string{}
	for _, v := range dfs(0) {
		ans = append(ans, strings.Join(v, " "))
	}
	return ans
}

// @lc code=end
