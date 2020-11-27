package problems

/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */
// 长度必相等，不等必false
// 构造26字母数组，对s字符串每个字符+1， t字符串每个字符-1
// 最后遍历26字母数组，都为0表示true
// @lc code=start
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	arr := [26]int{}

	for i, j := 0, 0; i < len(s); i, j = i+1, j+1 {
		arr[s[i]-'a']++
		arr[t[i]-'a']--
	}

	for _, s := range arr {
		if s != 0 {
			return false
		}
	}
	return true
}

// @lc code=end
