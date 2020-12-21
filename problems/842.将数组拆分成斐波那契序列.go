package problems

/*
 * @lc app=leetcode.cn id=842 lang=golang
 *
 * [842] 将数组拆分成斐波那契序列
 */

// @lc code=start
func SplitIntoFibonacci(S string) []int {
	n := len(S)
	ans := []int{}
	var backtrack func(idx, sum, prev int) bool
	backtrack = func(idx, sum, prev int) bool {
		if idx == n {
			return len(ans) >= 3
		}
		cur := 0
		for i := idx; i < n; i++ {
			if i > idx && S[idx] == '0' {
				break
			}
			cur = cur*10 + int(S[i]-'0')
			if cur > 1<<31-1 {
				break
			}

			if len(ans) >= 2 {
				if cur < sum {
					continue
				}
				if cur > sum {
					break
				}
			}
			ans = append(ans, cur)
			if backtrack(i+1, prev+cur, cur) {
				return true
			}
			ans = ans[:len(ans)-1]

		}
		return false
	}
	backtrack(0, 0, 0)
	return ans
}

// @lc code=end
