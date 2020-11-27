package problems

/*
 * @lc app=leetcode.cn id=409 lang=golang
 *
 * [409] 最长回文串
 */

// @lc code=start
func longestPalindrome(s string) int {
	m := make(map[int32]int)

	for _, v := range s {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	ans := 0
	for _, v := range m {
		ans += v / 2 * 2
	}
	if ans == len(s) {
		return ans
	}
	return ans + 1
}

// @lc code=end
