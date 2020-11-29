package problems

/*
 * @lc app=leetcode.cn id=367 lang=golang
 *
 * [367] 有效的完全平方数
 */

// @lc code=start
func isPerfectSquare(num int) bool {
	re := sqrt(num)
	return re == float64(int(re))
}

func sqrt(x int) float64 {
	if x == 0 {
		return 0
	}
	var last float64 = 0
	var res float64 = 1
	for res != last {
		last = res
		res = (res + float64(x)/res) / 2
	}
	return res
}

// @lc code=end
