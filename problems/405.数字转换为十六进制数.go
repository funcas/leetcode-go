package problems

/*
 * @lc app=leetcode.cn id=405 lang=golang
 *
 * [405] 数字转换为十六进制数
 */

// @lc code=start
func toHex(num int) string {
	if num == 0 {
		return "0"
	}

	c := []uint8{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	ans := []byte{}
	tmp := uint32(num)
	for tmp != 0 {
		ans = append([]byte{c[tmp&0b1111]}, ans...)
		tmp >>= 4
	}

	return string(ans)
}

// @lc code=end
