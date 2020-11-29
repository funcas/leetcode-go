package problems

/*
 * @lc app=leetcode.cn id=1002 lang=golang
 *
 * [1002] 查找常用字符
 */

// @lc code=start
func commonChars(A []string) []string {
	ans := make([]string, 0)
	c := [26]int{}
	for i := range c {
		c[i] = 1<<31 - 1
	}
	for _, word := range A {
		tmp := [26]int{}
		for _, b := range word {
			tmp[b-97]++
		}
		for i, f := range tmp {
			// 若比之前的小，说明重复个数少，取少
			if f < c[i] {
				c[i] = f
			}
		}
	}
	for i := 0; i < 26; i++ {
		for j := 0; j < c[i]; j++ {
			ans = append(ans, string(97+i))
		}
	}
	return ans
}

// @lc code=end

