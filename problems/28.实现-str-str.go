package problems

/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现 strStr()
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i] == needle[0] {
			j := 1
			for ; j < len(needle); j++ {
				if needle[j] != haystack[i+j] {
					break
				}
			}
			if j == len(needle) {
				return i
			}
		}
	}
	return -1
}

// @lc code=end
