package problems

/*
 * @lc app=leetcode.cn id=402 lang=golang
 *
 * [402] 移掉K位数字
 */

// @lc code=start
func removeKdigits(num string, k int) string {
	ans := "0"
	if k == len(num) {
		return ans
	}
	b := []byte(num)
	for i := 0; i < k; i++ {
		idx := 0
		for j := 1; j < len(b) && b[j] >= b[j-1]; j++ {
			idx = j
		}
		tmp := []byte{}
		tmp = append(tmp, b[:idx]...)
		tmp = append(tmp, b[idx+1:]...)
		b = tmp
	}
	for i := 0; i < len(b); i++ {
		if b[i] != '0' {
			ans = string(b[i:])
			break
		}

	}

	return ans
}

// RemoveKdigits2 利用栈加速查找过程
func removeKdigits2(num string, k int) string {
	var s []byte
	for i := 0; i < len(num); i++ {
		for k > 0 && len(s) > 0 && num[i] < s[len(s)-1] {
			k--
			s = s[:len(s)-1]
		}
		s = append(s, num[i])
	}
	s = s[:len(s)-k]
	ans := "0"
	for i := 0; i < len(s); i++ {
		if s[i] != '0' {
			ans = string(s[i:])
			break
		}
	}
	return ans
}

// @lc code=end
