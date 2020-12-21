package problems

/*
 * @lc app=leetcode.cn id=738 lang=golang
 *
 * [738] 单调递增的数字
 */

// @lc code=start
func monotoneIncreasingDigits(N int) int {
	i := 1
	ans := N
	for i <= ans/10 {
		n := ans / i % 100 // 取两位数字
		i *= 10
		if n/10 > n%10 { // 如果十位大于个位，取整-1为最大
			ans = ans/i*i - 1

		}
	}

	return ans
}

// @lc code=end
