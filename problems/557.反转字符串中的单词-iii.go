package problems

import "strings"

/*
 * @lc app=leetcode.cn id=557 lang=golang
 *
 * [557] 反转字符串中的单词 III
 */

// @lc code=start
func reverseWords(s string) string {
	arr := strings.Split(s, " ")
	for i := range arr {
		tmp := []byte(arr[i])
		for m, n := 0, len(tmp)-1;m < n;m,n=m+1,n-1 {
			tmp[m], tmp[n] = tmp[n], tmp[m]
		}
		arr[i] = string(tmp)
	}
	return strings.Join(arr, " ")
}
// @lc code=end

