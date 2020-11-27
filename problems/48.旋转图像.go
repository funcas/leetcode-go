package problems

/*
 * @lc app=leetcode.cn id=48 lang=golang
 *
 * [48] 旋转图像
 */

// @lc code=start
func rotate(matrix [][]int) {
	m := make(map[string]struct{})
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if i != j {
				k := string(i) + string(j)
				if _, ok := m[string(j)+string(i)]; !ok {
					matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
					m[k] = struct{}{}
				}

			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j, k := 0, len(matrix)-1; j < k; j, k = j+1, k-1 {
			matrix[i][j], matrix[i][k] = matrix[i][k], matrix[i][j]
		}
	}

}

// @lc code=end
