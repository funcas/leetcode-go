package problems

import "sort"

/*
 * @lc app=leetcode.cn id=1356 lang=golang
 *
 * [1356] 根据数字二进制下 1 的数目排序
 */

// @lc code=start
type item struct {
	Key int
	Val int
}

func sortByBits(arr []int) []int {
	l := make([]item, len(arr))
	ans := make([]int, len(arr))
	for i, v := range arr {
		tmp := item{
			Key: v,
			Val: count(v),
		}

		l[i] = tmp
	}
	sort.Slice(l, func(i, j int) bool {
		if l[i].Val == l[j].Val {
			return l[i].Key < l[j].Key
		}
		return l[i].Val < l[j].Val
	})
	for i, v := range l {
		ans[i] = v.Key
	}
	return ans
}

func count(num int) int {
	count := 0
	for num > 0 {
		num = num & (num - 1)
		count++
	}
	return count
}

// @lc code=end
