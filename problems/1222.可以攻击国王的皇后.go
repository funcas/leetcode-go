package problems

/*
 * @lc app=leetcode.cn id=1222 lang=golang
 *
 * [1222] 可以攻击国王的皇后
 */

// @lc code=start
func queensAttacktheKing(queens [][]int, king []int) [][]int {
	// row 2 col 2 pie 2 na 2
	arr := []int{-1, -1, -1, -1, -1, -1, -1, -1}
	kx, ky := king[0], king[1]
	for i, queen := range queens {
		qx, qy := queen[0], queen[1]
		if qx == kx {
			if qy < ky && (arr[0] < 0 || queens[arr[0]][1] < qy) {
				arr[0] = i
			}
			if qy > ky && (arr[1] < 0 || queens[arr[1]][1] > qy) {
				arr[1] = i
			}
		} else if qy == ky {
			if qx < kx && (arr[2] < 0 || queens[arr[2]][0] < qx) {
				arr[2] = i
			}
			if qx > kx && (arr[3] < 0 || queens[arr[3]][0] > qx) {
				arr[3] = i
			}
		} else if qx+qy == kx+ky {
			if qx < kx && (arr[4] < 0 || queens[arr[4]][0] < qx) {
				arr[4] = i
			}
			if qx > kx && (arr[5] < 0 || queens[arr[5]][0] > qx) {
				arr[5] = i
			}
		} else if qx-qy == kx-ky {
			if qx < kx && (arr[6] < 0 || queens[arr[6]][0] < qx) {
				arr[6] = i
			}
			if qx > kx && (arr[7] < 0 || queens[arr[7]][0] > qx) {
				arr[7] = i
			}
		}
	}
	ans := [][]int{}
	for _, v := range arr {
		if v != -1 {
			ans = append(ans, queens[v])
		}
	}
	return ans
}

// @lc code=end
