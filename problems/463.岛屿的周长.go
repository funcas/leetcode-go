package problems

/*
 * @lc app=leetcode.cn id=463 lang=golang
 *
 * [463] 岛屿的周长
 */

// @lc code=start
func islandPerimeter(grid [][]int) int {
	land := 0
	border := 0
	size := len(grid)
	for col := 0; col < size; col++ {

		for row := 0; row < len(grid[col]); row++ {
			if grid[col][row] == 1 {
				land++

				// 左
				if row-1 >= 0 && grid[col][row-1] == 1 {
					border++
				}
				// 上
				if col-1 >= 0 && grid[col-1][row] == 1 {
					border++
				}
			}

		}
	}
	return land*4 - border*2
}

// @lc code=end
