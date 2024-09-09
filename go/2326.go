package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}

	for i := range ans {
		for j := range ans[i] {
			ans[i][j] = -1
		}
	}

	t, b := 0, m-1
	l, r := 0, n-1

	for head != nil {
		for j := l; j < r+1; j++ {
			if head == nil {
				break
			}
			ans[t][j] = head.Val
			head = head.Next
		}
		for i := t + 1; i < b+1; i++ {
			if head == nil {
				break
			}
			ans[i][r] = head.Val
			head = head.Next
		}
		for j := r - 1; j > l-1; j-- {
			if head == nil {
				break
			}
			ans[b][j] = head.Val
			head = head.Next
		}
		for i := b - 1; i > t; i-- {
			if head == nil {
				break
			}
			ans[i][l] = head.Val
			head = head.Next
		}

		t++
		b--
		l++
		r--
	}

	return ans
}