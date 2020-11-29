package problems

/*
 * @lc app=leetcode.cn id=493 lang=golang
 *
 * [493] 翻转对
 */

// @lc code=start
func reversePairs(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var reverse func(left, right int) int
	reverse = func(left, right int) int {
		if left == right {
			return 0
		} else {
			mid := (left + right) / 2
			n1 := reverse(left, mid)
			n2 := reverse(mid+1, right)
			ret := n1 + n2

			i := left
			j := mid + 1
			for i <= mid {
				for j <= right && 2*nums[j] < nums[i] {
					j++
				}
				ret += j - mid - 1
				i++
			}

			sorted := make([]int, right-left+1)
			p1, p2 := left, mid+1
			p := 0
			for p1 <= mid || p2 <= right {
				if p1 > mid {
					sorted[p] = nums[p2]
					p, p2 = p+1, p2+1
				} else if p2 > right {
					sorted[p] = nums[p1]
					p, p1 = p+1, p1+1
				} else {
					if nums[p1] < nums[p2] {
						sorted[p] = nums[p1]
						p, p1 = p+1, p1+1
					} else {
						sorted[p] = nums[p2]
						p, p2 = p+1, p2+1
					}
				}

			}
			for k := range sorted {
				nums[left+k] = sorted[k]
			}

			return ret
		}
	}

	return reverse(0, len(nums)-1)
}

// @lc code=end
