package problems

/*
 * @lc app=leetcode.cn id=73 lang=golang
 *
 * [73] 矩阵置零
 */

// @lc code=start
func setZeroes(matrix [][]int) {
	xmap := make(map[int]struct{})
	ymap := make(map[int]struct{})
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				xmap[i] = struct{}{}
				ymap[j] = struct{}{}
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if _, ok := xmap[i]; ok {
				matrix[i][j] = 0
			}
			if _, ok := ymap[j]; ok {
				matrix[i][j] = 0
			}
		}
	}

}

// @lc code=end
