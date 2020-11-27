package problems

import "strings"

/*
 * @lc app=leetcode.cn id=71 lang=golang
 *
 * [71] 简化路径
 */

// @lc code=start
func simplifyPath(path string) string {
	if path == "" {
		return ""
	}
	t := []string{}
	p := strings.Split(path, "/")
	for _, q := range p {
		if q != "" && q != "." && q != ".." {
			t = append(t, q)
		}
		if q == ".." && len(t) > 0 {
			t = t[:len(t)-1]
		}
	}
	ans := "/" + strings.Join(t, "/")
	return ans
}

// @lc code=end
