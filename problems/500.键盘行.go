package problems

/*
 * @lc app=leetcode.cn id=500 lang=golang
 *
 * [500] 键盘行
 */

// @lc code=start
func findWords(words []string) []string {
	m := map[byte]int{'q': 1, 'w': 1, 'e': 1, 'r': 1, 't': 1, 'y': 1, 'u': 1, 'i': 1, 'o': 1, 'p': 1,
		'a': 2, 's': 2, 'd': 2, 'f': 2, 'g': 2, 'h': 2, 'j': 2, 'k': 2, 'l': 2,
		'z': 3, 'x': 3, 'c': 3, 'v': 3, 'b': 3, 'n': 3, 'm': 3}
	ans := []string{}
	for _, word := range words {
		tmp := 0
		for i := 0; i < len(word); i++ {
			c := word[i]
			if c >= 'A' && c <= 'Z' {
				c += 32
			}
			if tmp == 0 {
				tmp = m[c]
			} else if tmp != m[c] {
				tmp = -1
				break
			}
		}
		if tmp >= 0 {
			ans = append(ans, word)
		}
	}
	return ans
}

// @lc code=end
