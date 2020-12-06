package problems

/*
 * @lc app=leetcode.cn id=303 lang=golang
 *
 * [303] 区域和检索 - 数组不可变
 */
// a+b+c+d+...n = sum[n]
// c+d+...+n = sum[n] - (a+b)
// @lc code=start
type NumArray struct {
	sums []int
}

func Constructor(nums []int) NumArray {
	sums := make([]int, len(nums))
	if len(nums) == 0 {
		return NumArray{}
	}
	sums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sums[i] = sums[i-1] + nums[i]
	}
	return NumArray{
		sums: sums,
	}

}

func (this *NumArray) SumRange(i int, j int) int {
	if this.sums == nil {
		return 0
	}
	if i == 0 {
		return this.sums[j]
	}
	return this.sums[j] - this.sums[i-1]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
// @lc code=end
