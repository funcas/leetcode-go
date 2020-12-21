package problems

/*
 * @lc app=leetcode.cn id=643 lang=golang
 *
 * [643] 子数组最大平均数 I
 */

// @lc code=start
func FindMaxAverage(nums []int, k int) float64 {
	n := len(nums)
	sum := make([]int, n)
	sum[0] = nums[0]
	for i := 1; i < n; i++ {
		sum[i] = nums[i] + sum[i-1]
	}
	total := sum[k-1]
	for i := k; i < n; i++ {
		tmp := sum[i] - sum[i-k]
		if total < tmp {
			total = tmp
		}
	}
	return float64(total) / float64(k)
}

// @lc code=end

/*
1,12,-5,-6,50,3
1 13 8 2 52 55

*/
