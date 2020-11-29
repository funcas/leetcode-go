package problems

/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {

	mStack := make([]string, 0)
	p := 0
	for i := 0; i < len(s); i++ {
		if len(mStack) > len(s)/2 {
			return false
		}
		if p > 0 && isCouple(mStack[p-1], string(s[i])) {
			mStack = mStack[0 : p-1]
			p--
		} else {
			mStack = append(mStack, string(s[i]))
			p++
		}
	}
	if len(mStack) == 0 {
		return true
	}
	return false
}

func isCouple(a string, b string) bool {

	if a == "{" && b == "}" || a == "[" && b == "]" ||
		a == "(" && b == ")" {
		return true
	}
	return false
}

// @lc code=end
