package problems

/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 */

// @lc code=start
func SpiralOrder(matrix [][]int) []int {
	ans := []int{}
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return ans
	}
	top, left, right, bottom := 0, 0, len(matrix[0])-1, len(matrix)-1

	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			ans = append(ans, matrix[top][i])
		}
		for i := top + 1; i <= bottom; i++ {
			ans = append(ans, matrix[i][right])
		}
		if left < right && top < bottom {
			for i := right - 1; i > left; i-- {
				ans = append(ans, matrix[bottom][i])
			}
			for i := bottom; i > top; i-- {
				ans = append(ans, matrix[i][left])
			}
		}
		top++
		bottom--
		left++
		right--
	}
	return ans
}

// @lc code=end

/*

[ 1, 2, 3 ],
[ 4, 5, 6 ],
[ 7, 8, 9 ]


*/
