package problems

import (
	"math"
	"sort"
)

/*
 * @lc app=leetcode.cn id=1030 lang=golang
 *
 * [1030] 距离顺序排列矩阵单元格
 */

// @lc code=start
func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	m := make(map[int][][]int)
	ans := [][]int{}
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			dist := int(math.Abs(float64(i-r0)) + math.Abs(float64(j-c0)))
			if _, ok := m[dist]; ok {
				m[dist] = append(m[dist], []int{i, j})
			} else {
				m[dist] = [][]int{{i, j}}
			}
		}
	}
	keys := []int{}
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for k := range keys {
		ans = append(ans, m[k]...)
	}
	return ans
}

// @lc code=end
