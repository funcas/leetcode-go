package problems

/*
 * @lc app=leetcode.cn id=541 lang=golang
 *
 * [541] 反转字符串 II
 */

// @lc code=start
func ReverseStr(s string, k int) string {
	n := len(s)
	b := []byte(s)
	if n < k {
		for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
			b[i], b[j] = b[j], b[i]
		}
		return string(b)
	}

	for i := 0; 2*k*i < n; i++ {
		begin := 2 * k * i
		end := 2*k*i + k - 1
		if end >= n {
			end = n - 1
		}
		for j, k := begin, end; j < k; j, k = j+1, k-1 {
			b[j], b[k] = b[k], b[j]
		}
	}
	return string(b)
}

// @lc code=end
