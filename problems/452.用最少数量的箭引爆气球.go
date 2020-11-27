package problems

import "sort"

/*
 * @lc app=leetcode.cn id=452 lang=golang
 *
 * [452] 用最少数量的箭引爆气球


        标杆
[1.......6]
   [2..........8]
            [7.........12]
                   [10.........16]

拿当前区间的右端作为标杆。
只要 下一个区间的左端<=标杆下一个区间的左端<=标杆，则重合，继续寻求与下一个区间重合。
直到遇到不重合的区间，弓箭数 +1。
拿新区间的右端作为标杆，重复以上步骤。
*/

// @lc code=start
func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	count := 1
	x := points[0][1]
	for i := 1; i < len(points); i++ {
		if x < points[i][0] {
			count++
			x = points[i][1]
		}
	}
	return count

}

// @lc code=end
