package problems

/*
 * @lc app=leetcode.cn id=1207 lang=golang
 *
 * [1207] 独一无二的出现次数
 */

// @lc code=start
func uniqueOccurrences(arr []int) bool {
	m := make(map[int]int)
	for _, it := range arr {
		if v, ok := m[it]; ok {
			m[it] = v + 1
		} else {
			m[it] = 1
		}
	}
	n := make(map[int]struct{})

	for _, v := range m {
		n[v] = struct{}{}
	}
	return len(m) == len(n)
}

// @lc code=end
