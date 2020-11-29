package problems

import "fmt"

/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
func isPalindrome(x int) bool {
	t := fmt.Sprintf("%d", x)
	for i, j := 0, len(t)-1; i <= j; i, j = i+1, j-1 {
		if t[i] != t[j] {
			return false
		}
	}
	return true
}

// 不转字符串方式
func isPalindrome2(x int) bool {

	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	a := 0
	for x > a {
		a = a*10 + x%10
		x = x / 10
	}
	return a == x || x == a/10

}

func IsPalindrome3(s string) bool {
	tmp := []int32{}
	for _, v := range s {
		if (v >= '0' && v <= '9') || (v >= 'a' && v <= 'z') {

			tmp = append(tmp, v)
		}
		if v >= 'A' && v <= 'Z' {
			tmp = append(tmp, v+32)
		}

	}
	for i, j := 0, len(tmp)-1; i < j; i, j = i+1, j-1 {
		if tmp[i] != tmp[j] {
			return false
		}
	}
	return true

}

// @lc code=end
