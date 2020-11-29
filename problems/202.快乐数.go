package problems

/*
 * @lc app=leetcode.cn id=202 lang=golang
 *
 * [202] 快乐数
 */

// @lc code=start
func isHappy(n int) bool {
	m := make(map[int]struct{})
	var happy func(n int) bool

	happy = func(n int) bool {
		ret := getSquare(n)
		if ret == 1 {
			return true
		}
		if _, ok := m[ret]; ok {
			return false
		}
		m[ret] = struct{}{}
		return happy(ret)
	}

	return happy(n)
}

func getSquare(n int) int {
	t := n
	total := 0
	for t > 0 {
		x := t % 10
		t = t / 10
		total += x * x
	}
	return total
}

// @lc code=end
