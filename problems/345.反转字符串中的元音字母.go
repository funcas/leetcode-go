package problems

/*
 * @lc app=leetcode.cn id=345 lang=golang
 *
 * [345] 反转字符串中的元音字母
 */

// @lc code=start
func reverseVowels(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; {
		for i < j && !isVowel(b[i]) {
			i++
		}
		for i < j && !isVowel(b[j]) {
			j--
		}
		if i < j {
			b[i], b[j] = b[j], b[i]
		}
		i++
		j--
	}
	return string(b)
}

func isVowel(s uint8) bool {
	return s == 'a' || s == 'e' || s == 'i' || s == 'o' || s == 'u' ||
		s == 'A' || s == 'E' || s == 'I' || s == 'O' || s == 'U'
}

// @lc code=end
