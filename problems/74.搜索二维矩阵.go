package problems

/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */

// @lc code=start
func SearchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	m, n := len(matrix)*len(matrix[0]), len(matrix[0])
	begin, end := 0, m

	for begin <= end {
		mid := (begin + end) / 2
		if mid/n >= len(matrix) {
			break
		}
		item := matrix[mid/n][mid%n]
		if item == target {
			return true
		}
		if item < target {
			begin = mid + 1
		}

		if item > target {
			end = mid - 1
		}
	}
	return false
}

// @lc code=end
