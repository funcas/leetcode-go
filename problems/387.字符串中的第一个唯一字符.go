package problems

/*
 * @lc app=leetcode.cn id=387 lang=golang
 *
 * [387] 字符串中的第一个唯一字符
 */

// @lc code=start
func firstUniqChar(s string) int {
	arr := make([]int, 26)

	for _, v := range s {
		arr[v-'a']++
	}
	ans := -1
	for i, v := range s {
		if arr[v-'a'] == 1 {
			ans = i
			break
		}
	}

	return ans
}

// @lc code=end
