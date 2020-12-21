package problems

/*
 * @lc app=leetcode.cn id=482 lang=golang
 *
 * [482] 密钥格式化
 */

// @lc code=start
func licenseKeyFormatting(S string, K int) string {
	st := []byte{}
	for i := 0; i < len(S); i++ {
		if S[i] != '-' {
			if S[i] >= 'a' && S[i] <= 'z' {
				st = append(st, S[i]-32)
			} else {
				st = append(st, S[i])
			}

		}
	}
	ans := []byte{}
	cnt := 0
	for i := len(st) - 1; i >= 0; i-- {
		ans = append(ans, st[i])
		cnt++
		if cnt%K == 0 && i != 0 {
			ans = append(ans, '-')
		}
	}
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return string(ans)
}

// @lc code=end
