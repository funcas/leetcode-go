package problems

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=621 lang=golang
 *
 * [621] 任务调度器

   执行顺序的规律是： 有count - 1个A，其中每个A需要搭配n个X，再加上最后一个A，所以总时间为 (count - 1) * (n + 1) + 1
   要注意可能会出现多个频率相同且都是最高的任务，比如 ["A","A","A","B","B","B"]，
   所以最后会剩下一个A和一个B，因此最后要加上频率最高的不同任务的个数 maxCount
*/

// @lc code=start
func leastInterval(tasks []byte, n int) int {
	arr := make([]int, 26)
	for _, v := range tasks {
		arr[v-'A']++
	}
	sort.Ints(arr)
	maxCnt := 0

	for i := 25; i >= 0; i-- {
		if arr[i] != arr[25] {
			break
		}
		maxCnt++
	}

	return max((arr[25]-1)*(n+1)+maxCnt, len(tasks))
}

// @lc code=end
