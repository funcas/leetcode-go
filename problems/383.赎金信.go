package problems

/*
 * @lc app=leetcode.cn id=383 lang=golang
 *
 * [383] 赎金信
 */

// @lc code=start
func canConstruct(ransomNote string, magazine string) bool {
	if len(magazine) < len(ransomNote) {
		return false
	}
	arr := make([]int, 26)

	for i := 0; i < len(magazine); i++ {
		if i < len(ransomNote) {
			arr[ransomNote[i]-'a']--
		}
		arr[magazine[i]-'a']++
	}
	for _, v := range arr {
		if v < 0 {
			return false
		}
	}
	return true
}

// @lc code=end
