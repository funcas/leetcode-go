package problems

/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	m := make(map[uint8]int)
	ans := 0
	for i, j := 0, 0; i < len(s); i++ {
		if m[s[i]] != 0 {
			j = max(j, int(m[s[i]]))
		}
		ans = max(i-j+1, ans)
		m[s[i]] = i + 1
	}
	return ans
}

// @lc code=end
