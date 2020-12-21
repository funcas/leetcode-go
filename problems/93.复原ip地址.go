package problems

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=93 lang=golang
 *
 * [93] 复原IP地址
 */

// @lc code=start
func RestoreIpAddresses(s string) []string {
	n := len(s)
	ans := []string{}
	if n > 12 {
		return []string{}
	}
	result := [][]string{}
	var backtrack func(idx int, cnt int, path []string)
	backtrack = func(idx int, cnt int, path []string) {
		if len(path) == 4 && cnt < n {
			return
		}

		if len(path) == 4 {
			tmp := make([]string, 4)
			copy(tmp, path)
			result = append(result, tmp)
			return
		}

		for i := 0; i < 3; i++ {
			if idx+i+1 > n {
				continue
			}
			p := s[idx : idx+i+1]
			pint, _ := strconv.Atoi(p)
			if (len(p) > 1 && p[0] == '0') || pint > 255 {
				continue
			}
			path = append(path, s[idx:idx+i+1])
			backtrack(idx+i+1, cnt+i+1, path)
			path = path[:len(path)-1]
		}

	}
	backtrack(0, 0, []string{})
	for _, v := range result {
		str := strings.Join(v, ".")
		ans = append(ans, str)
	}
	return ans
}

// @lc code=end
