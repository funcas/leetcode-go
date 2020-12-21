package problems
/*
 * @lc app=leetcode.cn id=605 lang=golang
 *
 * [605] 种花问题
 */

// @lc code=start
func canPlaceFlowers(flowerbed []int, n int) bool {
	flowerbed = append(append([]int{0}, flowerbed...), 0)
	for i:=1;i<len(flowerbed)-1;i++{
		if flowerbed[i-1]==0 && flowerbed[i]==0 && flowerbed[i+1]==0 {
			flowerbed[i]= 1
			n -=1
		}
	}
	return n <= 0

}
// @lc code=end

