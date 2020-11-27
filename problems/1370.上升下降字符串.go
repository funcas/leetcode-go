package problems

import (
	"bytes"
)

/*
 * @lc app=leetcode.cn id=1370 lang=golang
 *
 * [1370] 上升下降字符串
 */

// @lc code=start
func sortString(s string) string {
	var b bytes.Buffer
	arr := [26]int{}
	for i := 0; i < len(s); i++ {
		arr[s[i]-'a']++
	}
	i := 0
	for b.Len() < len(s) {
		for ; i < 26; i++ {
			if arr[i] > 0 {
				b.WriteString(string(uint8(i) + 'a'))
				arr[i]--
			}
		}
		i = 25
		for ; i >= 0; i-- {
			if arr[i] > 0 {
				b.WriteString(string(uint8(i) + 'a'))
				arr[i]--
			}
		}
		i = 0
	}
	return b.String()
}

// @lc code=end
