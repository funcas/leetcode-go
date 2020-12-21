package problems

/*
 * @lc app=leetcode.cn id=594 lang=golang
 *
 * [594] 最长和谐子序列
 */

// @lc code=start
func findLHS(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	ans := 0
	for _, v := range nums {
		t := max(m[v-1], m[v+1])
		if t != 0 {
			ans = max(ans, m[v]+t)
		}
	}
	return ans
}

// @lc code=end
