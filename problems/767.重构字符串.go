package problems

/*
 * @lc app=leetcode.cn id=767 lang=golang
 *
 * [767] 重构字符串
 */

// @lc code=start
func ReorganizeString(S string) string {
	size := len(S)
	if size <= 1 {
		return S
	}
	maxCh := 0
	arr := [26]int{}
	for _, v := range S {
		arr[v-'a']++
		if arr[v-'a'] > maxCh {
			maxCh = arr[v-'a']
		}
	}
	if maxCh > (1+size)/2 {
		return ""
	}
	ans := make([]byte, size)
	even, odd, half := 0, 1, size/2
	for i := 0; i < 26; i++ {
		v := arr[i]
		tmp := byte(i + 'a')
		for v > 0 && v <= half && odd < size {
			ans[odd] = tmp
			v--
			odd += 2
		}
		for v > 0 {
			ans[even] = tmp
			v--
			even += 2
		}
	}
	return string(ans)
}

// @lc code=end
