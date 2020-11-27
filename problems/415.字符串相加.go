package problems

import "bytes"

/*
 * @lc app=leetcode.cn id=415 lang=golang
 *
 * [415] 字符串相加
 */

// @lc code=start
func addStrings(num1 string, num2 string) string {
	var b bytes.Buffer

	i, j := len(num1)-1, len(num2)-1
	carry := byte(0)
	for i >= 0 || j >= 0 || carry > 0 {
		if i >= 0 {
			carry += num1[i] - '0'
		}
		if j >= 0 {
			carry += num2[j] - '0'
		}
		b.WriteString(string(carry%10 + '0'))
		carry /= 10
		i--
		j--
	}
	ans := b.Bytes()
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return string(ans)
}

// @lc code=end
