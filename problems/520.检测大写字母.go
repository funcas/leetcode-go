package problems

/*
 * @lc app=leetcode.cn id=520 lang=golang
 *
 * [520] 检测大写字母
 */

// @lc code=start
func detectCapitalUse(word string) bool {

	n := len(word)
	count := 0
	for i := 0; i < n; i++ {
		if word[i] >= 65 && word[i] <= 90 {
			count++
		}
	}
	if n == count {
		return true
	} else if count == 0 {
		return true
	} else if count == 1 && (word[0] >= 65 && word[0] <= 90) {
		return true
	} else {
		return false
	}

}

// @lc code=end
// USA
