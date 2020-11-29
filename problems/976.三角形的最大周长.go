package problems

import "sort"

/*
 * @lc app=leetcode.cn id=976 lang=golang
 *
 * [976] 三角形的最大周长
 */
// a <= b <= c && a + b < c
// 排序，逆序遍历，第一组满足条件的即为解
// @lc code=start
func largestPerimeter(A []int) int {
	sort.Ints(A)

	for i := len(A) - 1; i >= 2; i-- {
		a := A[i]
		b := A[i-1]
		c := A[i-2]
		if a < b+c {
			return a + b + c
		}
	}
	return 0
}

// @lc code=end
