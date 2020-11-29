package problems

/*
 * @lc app=leetcode.cn id=844 lang=golang
 *
 * [844] 比较含退格的字符串
 */

// @lc code=start
func backspaceCompare(S string, T string) bool {
	s := make([]string, 0)
	t := make([]string, 0)

	for _, ch := range S {
		if ch == '#' {
			if len(s) > 0 {
				s = s[:len(s)-1]
			}
			continue
		}
		s = append(s, string(ch))
	}

	for _, ch := range T {
		if ch == '#' {
			if len(t) > 0 {
				t = t[:len(t)-1]
			}
			continue
		}
		t = append(t, string(ch))
	}
	if len(s) == len(t) {
		for i := 0; i < len(s); i++ {
			if s[i] != t[i] {
				return false
			}
		}
		return true
	}
	return false
}

// @lc code=end
