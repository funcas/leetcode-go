package problems

import (
	"sort"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=506 lang=golang
 *
 * [506] 相对名次
 */

// @lc code=start
func findRelativeRanks(nums []int) []string {
	m := make(map[int]int)
	for i, v := range nums {
		m[v] = i
	}
	keys := []int{}
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] > keys[j]
	})
	arr := [...]string{"Gold Medal", "Silver Medal", "Bronze Medal"}
	ans := make([]string, len(nums))
	for i, v := range keys {
		if i < 3 {
			ans[m[v]] = arr[i]
		} else {
			ans[m[v]] = strconv.Itoa(i + 1)
		}

	}
	return ans
}

// @lc code=end

// 0
// 1 4 3 2 2 1 9
// 1 3 1 9
