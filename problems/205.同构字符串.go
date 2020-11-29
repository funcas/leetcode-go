package problems

/*
 * @lc app=leetcode.cn id=205 lang=golang
 *
 * [205] 同构字符串
 */

// @lc code=start
func isIsomorphic(s string, t string) bool {
	k1 := make(map[uint8]uint8)
	k2 := make(map[uint8]uint8)
	for i := 0; i < len(s); i++ {
		if r, ok := k1[s[i]]; ok {
			if r != t[i] {
				return false
			}
		}
		if q, ok := k2[t[i]]; ok {
			if q != s[i] {
				return false
			}
		}
		k1[s[i]] = t[i]
		k2[t[i]] = s[i]

	}
	return true
}

// @lc code=end
