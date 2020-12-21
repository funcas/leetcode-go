package problems

/*
 * @lc app=leetcode.cn id=91 lang=golang
 *
 * [91] 解码方法
 */
// 当前位为0，如果前一位为1或2，则只能是当前位和前一位一起解码，dp[i] = dp[i-2]，否则，直接返回0；
// 前一位为1，且当前为不为0，则有当前位单独解码和与前一位一起解码两种情况，即dp[i] = dp[i-1] + dp[i-2];
// 前一位为2，且当前为在1~6之间，则有当前位单独解码和与前一位一起解码两种情况，即dp[i] = dp[i-1] + dp[i-2];
// 其他情况，只能是单独解码，dp[i] = dp[i-1]。
// @lc code=start
func NumDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	n := len(s)
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				if i == 1 {
					dp[i] = 1
				} else {
					dp[i] = dp[i-2]
				}
			} else {
				return 0
			}
		} else if s[i-1] == '1' || (s[i-1] == '2' && s[i] >= '1' && s[i] <= '6') {
			if i == 1 {
				dp[i] = 2
			} else {
				dp[i] = dp[i-1] + dp[i-2]
			}
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[n-1]
}

// @lc code=end
// 2 2 1 4 8

// 空间优化
func NumDecodings3(s string) int {
	if s[0] == '0' {
		return 0
	}
	n := len(s)
	prev, cur := 1, 1 //dp[-1]=dp[0] = 1
	for i := 1; i < n; i++ {
		tmp := cur
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				cur = prev
			} else {
				return 0
			}
		} else if s[i-1] == '1' || (s[i-1] == '2' && s[i] >= '1' && s[i] <= '6') {
			cur = cur + prev
		}
		prev = tmp
	}
	return cur
}

// 回溯，超时
func numDecodings2(s string) int {
	n := len(s)
	ans := 0
	var backtrack func(idx int, path []int)
	backtrack = func(idx int, path []int) {

		if idx == n {
			ans++
			// fmt.Println(path)
			return
		}
		cur := 0
		for i := idx; i < n; i++ {
			if i > idx && s[idx] == '0' {
				break
			}
			cur = cur*10 + int(s[i]-'0')
			if cur == 0 || cur > 26 {
				break
			}
			path = append(path, cur)
			backtrack(i+1, path)
			path = path[:len(path)-1]
		}

	}
	backtrack(0, []int{})
	return ans
}

// 1 1 8 2
// dp0 = 1    -- 1
// dp1 = 2    -- 11 1,1
// dp2 = 3    -- 1,1,8 11,8 1,18
// dp3 = 3     -- 1,1,8,2 11,8,2 1,18,2
