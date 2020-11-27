package problems

/*
 * @lc app=leetcode.cn id=89 lang=golang
 *
 * [89] 格雷编码
 */

// @lc code=start
func grayCode(n int) []int {
	ans := []int{}
	if n == 0 {
		return []int{0}
	}

	m := make(map[int]struct{})
	var calc func(num int)
	calc = func(num int) {
		// if _, ok := m[num]; ok {
		// 	return
		// }
		m[num] = struct{}{}
		ans = append(ans, num)
		for i := 0; i < n; i++ {
			num := num ^ (1 << i)
			if _, ok := m[num]; ok {
				continue
			}
			calc(num)
		}
	}
	calc(0)

	return ans
}

func grayCode2(n int) []int {
	c := 1 << n
	ans := []int{}
	for i := 0; i < c; i++ {
		ans = append(ans, (i>>1)^i)
	}
	return ans
}

// @lc code=end
