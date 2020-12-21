package problems

/*
 * @lc app=leetcode.cn id=744 lang=golang
 *
 * [744] 寻找比目标字母大的最小字母
 */

// @lc code=start
func NextGreatestLetter(letters []byte, target byte) byte {
	n := len(letters)
	i, j := 0, n
	for i < j {
		mid := (i + j) / 2
		if letters[mid] <= target {
			i = mid + 1
		} else {
			j = mid
		}
	}
	return letters[i%n]
}

// @lc code=end
