package problems

import "fmt"

/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 */

// @lc code=start
func LargestRectangleArea(heights []int) int {
	s := []int{}
	size := len(heights)
	l, r := make([]int, size), make([]int, size)
	for i := 0; i < size; i++ {
		for len(s) > 0 && heights[i] <= heights[s[len(s)-1]] {
			s = s[:len(s)-1]
		}
		if len(s) == 0 {
			l[i] = -1
		} else {
			l[i] = s[len(s)-1]
		}
		s = append(s, i)
	}
	s = []int{}
	for i := size - 1; i >= 0; i-- {
		for len(s) > 0 && heights[i] <= heights[s[len(s)-1]] {
			s = s[:len(s)-1]
		}
		if len(s) == 0 {
			r[i] = size
		} else {
			r[i] = s[len(s)-1]
		}
		s = append(s, i)
	}
	fmt.Println(l)
	fmt.Println(r)
	ans := 0
	for i := 0; i < size; i++ {
		ans = max(ans, (r[i]-l[i]-1)*heights[i])
	}
	return ans
}

// @lc code=end

// 2,1,5,6,2,3
// -1 -1 2 3 1 4
