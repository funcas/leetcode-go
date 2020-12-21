package problems

import "regexp"

/*
 * @lc app=leetcode.cn id=551 lang=golang
 *
 * [551] 学生出勤记录 I
 */

// @lc code=start
func checkRecord(s string) bool {
	match, _ := regexp.MatchString(".*(A.*A|LLL).*", s)
	return !match
}

// @lc code=end

// PPALLP
