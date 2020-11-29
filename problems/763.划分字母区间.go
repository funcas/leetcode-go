package problems

/*
 * @lc app=leetcode.cn id=763 lang=golang
 *
 * [763] 划分字母区间
 */

// @lc code=start
func partitionLabels(S string) []int {
	ans := []int{}
	lastPos := [26]int{}
	for i, c := range S {
		lastPos[c-97] = i
	}
	start, end := 0, 0
	for i, c := range S {
		if lastPos[c-97] > end {
			end = lastPos[c-97]
		}
		if i == end {
			ans = append(ans, end-start+1)
			start = end + 1
		}
	}
	return ans
}

// @lc code=end
