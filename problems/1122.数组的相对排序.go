package problems

import "sort"

/*
 * @lc app=leetcode.cn id=1122 lang=golang
 *
 * [1122] 数组的相对排序
 */

// @lc code=start
func relativeSortArray(arr1 []int, arr2 []int) []int {
	m := make(map[int]int)
	for _, v := range arr1 {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	ans := []int{}
	for i := range arr2 {
		item := arr2[i]
		for j := 0; j < m[item]; j++ {
			ans = append(ans, item)
		}
		delete(m, item)
	}

	if len(m) > 0 {
		remain := []int{}
		for k, v := range m {
			for i := 0; i < v; i++ {
				remain = append(remain, k)
			}
		}
		sort.Ints(remain)
		ans = append(ans, remain...)
	}
	return ans
}

// @lc code=end
