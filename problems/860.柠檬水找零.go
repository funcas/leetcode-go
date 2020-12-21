package problems

/*
 * @lc app=leetcode.cn id=860 lang=golang
 *
 * [860] 柠檬水找零
 */

// @lc code=start
func LemonadeChange(bills []int) bool {
	if bills[0] != 5 {
		return false
	}
	o := make([]int, 2)
	for _, val := range bills {
		if val == 5 {
			o[0]++
		} else if val == 10 {
			if o[0] > 0 {
				o[0]--
				o[1]++
			} else {
				return false
			}
		} else if val == 20 {
			if o[1] > 0 && o[0] > 0 {
				o[1]--
				o[0]--
			} else if o[1] == 0 && o[0] >= 3 {
				o[0] -= 3
			} else {
				return false
			}
		}
	}
	return true
}

// @lc code=end
