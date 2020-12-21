package problems

/*
 * @lc app=leetcode.cn id=12 lang=golang
 *
 * [12] 整数转罗马数字
 */

// @lc code=start
func intToRoman(num int) string {
	dig := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	ten := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	hun := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	thu := []string{"", "M", "MM", "MMM"}
	return thu[num/1000] + hun[num%1000/100] + ten[num%100/10] + dig[num%10]
}

// @lc code=end
