package problems

import (
	"bytes"
)

/*
 * @lc app=leetcode.cn id=504 lang=golang
 *
 * [504] 七进制数
 */

// @lc code=start
func ConvertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	var b bytes.Buffer
	flag := false
	if num < 0 {
		flag = true
		num = -num
	}
	for num != 0 {
		b.WriteByte(uint8(num%7) + '0')
		num = num / 7
	}
	tmp := []byte(b.String())
	for i, j := 0, len(tmp)-1; i < j; i, j = i+1, j-1 {
		tmp[i], tmp[j] = tmp[j], tmp[i]
	}
	if flag {
		return "-" + string(tmp)
	}
	return string(tmp)
}

// @lc code=end
