package problems

/*
 * @lc app=leetcode.cn id=520 lang=golang
 *
 * [520] 检测大写字母
 */

// @lc code=start
func detectCapitalUse(word string) bool {

	for i := 2; i < len(word); i++ {
		if word[0] <= 'Z' && word[1] <= 'Z' {
			if word[i] > 'Z' {
				return false
			}
		} else if word[0] <= 'Z' && word[1] > 'Z' {
			if word[i] <= 'Z' {
				return false
			}
		} else if word[0] > 'Z' && word[1] > 'Z' {
			if word[i] <= 'Z' {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

// @lc code=end
// USA
