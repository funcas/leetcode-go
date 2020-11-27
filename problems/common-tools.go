package problems

type ListNode struct {
	Val  int
	Next *ListNode
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 取大值
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 取小值
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 交换
func swap(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

// 快排
func qc(arr []int, p int, r int) {
	if p >= r {
		return
	}
	q := partation(arr, p, r)
	qc(arr, p, q)
	qc(arr, q+1, r)
}

func partation(arr []int, p int, r int) int {
	j := p
	provit := arr[r-1]

	for i := p; i < r-1; i++ {
		if arr[i] < provit {
			swap(&arr[j], &arr[i])
			j++
		}
	}
	arr[j], arr[r-1] = arr[r-1], arr[j]
	return j
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
