package problems

import (
	"math"
)

/*
 * @lc app=leetcode.cn id=492 lang=golang
 *
 * [492] 构造矩形
 */

// @lc code=start
func ConstructRectangle(area int) []int {
	mid := int(math.Sqrt(float64(area)))
	L := mid
	for ; L > 0; L-- {
		if area%L == 0 {
			break
		}
	}
	return []int{L, area / L}
}

// @lc code=end
