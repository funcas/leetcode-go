package problems

import (
	"math"
)

/*
 * @lc app=leetcode.cn id=507 lang=golang
 *
 * [507] 完美数
 */

// @lc code=start
func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}
	sum := 1
	n := int(math.Sqrt(float64(num)))
	for i := 2; i <= n; i++ {
		a := num % i
		b := i * i
		if a == 0 && b != num {
			sum += i
			sum += num / i
		}
		if b == num {
			sum += i
		}

	}
	return sum == num
}

// @lc code=end
