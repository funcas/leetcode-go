package problems

/*
 * @lc app=leetcode.cn id=454 lang=golang
 *
 * [454] 四数相加 II
 */

// @lc code=start
func FourSumCount(A []int, B []int, C []int, D []int) int {
	ans := 0
	m := make(map[int]int)
	for _, v := range A {
		for _, vb := range B {
			m[v+vb]++
		}
	}
	for _, vc := range C {
		for _, vd := range D {
			if x, ok := m[0-(vc+vd)]; ok {
				ans += x
			}
		}
	}
	return ans
}

// @lc code=end
