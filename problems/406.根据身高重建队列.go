package problems

import "sort"

/*
 * @lc app=leetcode.cn id=406 lang=golang
 *
 * [406] 根据身高重建队列
 */

// @lc code=start
func reconstructQueue(people [][]int) [][]int {
	ans := make([][]int, len(people))

	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] > people[j][1]
		}
		return people[i][0] < people[j][0]
	})

	for _, v := range people {
		t := 0
		for i := 0; i < len(ans); i++ {
			if ans[i] != nil {
				continue
			}
			if v[1] == t {
				ans[i] = v
				break
			}
			t++
		}
	}
	return ans
}

// @lc code=end
