package problems

/*
 * @lc app=leetcode.cn id=389 lang=golang
 *
 * [389] 找不同
 */

// @lc code=start
func findTheDifference(s string, t string) byte {
	arr := make([]int, 26)

	for i := 0; i < len(t); i++ {
		if i < len(s) {
			arr[s[i]-'a']++
		}
		arr[t[i]-'a']--
	}
	for i, v := range arr {
		if v < 0 {
			return byte(i + 'a')
		}
	}
	return 'a'
}

// @lc code=end
