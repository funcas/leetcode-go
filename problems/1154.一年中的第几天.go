package problems

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=1154 lang=golang
 *
 * [1154] 一年中的第几天
 */

// @lc code=start
func DayOfYear(date string) int {
	m := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	d := strings.Split(date, "-")
	flag := false
	ans := 0
	year, _ := strconv.Atoi(d[0])
	mon, _ := strconv.Atoi(d[1])
	day, _ := strconv.Atoi(d[2])

	if year%4 == 0 {
		if year%100 == 0 {
			flag = year%400 == 0
		} else {
			flag = true
		}

	}
	for i := 0; i < mon-1; i++ {
		ans += m[i]
	}
	ans += day
	if flag && mon > 2 {
		ans++
	}

	return ans

}

// @lc code=end
