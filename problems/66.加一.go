package problems

/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {
	res := []int{0}
	res = append(res, digits...)
	for i := len(res) - 1; i >= 0; i-- {
		x := res[i] + 1
		if x == 10 {
			res[i] = 0
		} else {
			res[i] = x
			break
		}
	}
	if res[0] == 0 {
		res = res[1:]
	}
	return res
}

// @lc code=end
