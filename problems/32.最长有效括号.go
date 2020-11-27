package problems

/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 */

// @lc code=start
type Bracket struct {
	Idx int
	Ch  uint8
}

func longestValidParentheses(s string) int {
	stack := []*Bracket{}
	arr := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && (s[i] == ')' && stack[len(stack)-1].Ch == '(') {
			tmp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			arr[tmp.Idx] = 1
			arr[i] = 1
		} else {
			stack = append(stack, &Bracket{
				Ch:  s[i],
				Idx: i,
			})
		}
	}

	ans := 0
	cnt := 0
	for _, v := range arr {
		if v == 1 {
			cnt++
		} else {
			if cnt > ans {
				ans = cnt
			}
			cnt = 0
		}
	}
	if cnt > ans {
		ans = cnt
	}
	return ans
}

// @lc code=end
