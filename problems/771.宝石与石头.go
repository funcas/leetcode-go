package problems

/*
 * @lc app=leetcode.cn id=771 lang=golang
 *
 * [771] 宝石与石头
 */

// @lc code=start
func numJewelsInStones(J string, S string) int {
	m := make(map[int32]int)
	for _, s := range S {
		if _, ok := m[s]; ok {
			m[s] += 1
		} else {
			m[s] = 1
		}
	}
	ans := 0
	for _, j := range J {
		if _, ok := m[j]; ok {
			ans += m[j]
		}
	}
	return ans
}

// @lc code=end
