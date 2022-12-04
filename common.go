package algo

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(nums []interface{}) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	ton := func(i interface{}) int {
		sprintf := fmt.Sprintf("%v", i)
		rv, _ := strconv.Atoi(sprintf)
		return rv
	}

	root := &TreeNode{Val: ton(nums[0])}
	q := []*TreeNode{root}
	i := 1
	for i < len(nums) {
		n := len(q)
		for j := 0; j < n; j++ {
			t := q[0]
			q = q[1:]
			for k := 0; k < 2 && i < len(nums); k++ {
				v := nums[i]
				var nn *TreeNode
				if v != "null" {
					nn = &TreeNode{Val: ton(v)}
					q = append(q, nn)
				}
				if k == 0 {
					t.Left = nn
				} else {
					t.Right = nn
				}
				i++
			}
		}
	}
	return root
}

func buildList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{nums[0], nil}
	cur := head
	for i := 1; i < len(nums); i++ {
		n := &ListNode{nums[i], nil}
		cur.Next = n
		cur = n
	}
	return head
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func fill1(array []int, n int) {
	for i := 0; i < len(array); i++ {
		array[i] = n
	}
}

func fill2(array [][]int, n int) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			array[i][j] = n
		}
	}
}

func initArray(m, n int) [][]int {
	ret := make([][]int, m)
	for i := 0; i < len(ret); i++ {
		ret[i] = make([]int, n)
	}
	return ret
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
