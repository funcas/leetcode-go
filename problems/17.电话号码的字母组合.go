package problems

/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
func letterCombinations(digits string) []string {
	dict := map[int32]string{'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno",
		'7': "pqrs", '8': "tuv", '9': "wxyz"}
	ans := []string{}
	for _, s := range digits {
		if v, ok := dict[s]; ok {
			if len(ans) == 0 {
				for i := 0; i < len(v); i++ {
					ans = append(ans, string(v[i]))
				}
				continue
			}
			tmp := []string{}
			for i := 0; i < len(ans); i++ {
				for j := 0; j < len(v); j++ {
					tmp = append(tmp, ans[i]+string(v[j]))
				}
			}
			ans = tmp
		}
	}
	return ans
}

// @lc code=end
