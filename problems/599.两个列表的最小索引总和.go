package problems

/*
 * @lc app=leetcode.cn id=599 lang=golang
 *
 * [599] 两个列表的最小索引总和
 */

// @lc code=start
func FindRestaurant(list1 []string, list2 []string) []string {
	ans := make(map[int][]string)
	m := make(map[string]int)
	for i, v := range list1 {
		m[v] = i
	}
	for i, v := range list2 {
		if j, ok := m[v]; ok {
			ans[i+j] = append(ans[i+j], v)
		}
	}
	tmp := 1<<32 - 1
	for k := range ans {
		if k < tmp {
			tmp = k
		}
	}
	return ans[tmp]
}

// @lc code=end
