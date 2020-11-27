package problems

/*
 * @lc app=leetcode.cn id=292 lang=golang
 *
 * [292] Nim 游戏
 */
// 找规律，4的倍数必输。。。淦
// win: 1 2 3 5 6 7 9 10 11 13 14 15 17 18 19 21 22 23 25 26 27
// def: 4 8 12 16 20 24 28
// @lc code=start
func canWinNim(n int) bool {
	return n%4 != 0
}

// @lc code=end
