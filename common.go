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
	var build func(nums []interface{}, index int) *TreeNode
	build = func(nums []interface{}, index int) *TreeNode {
		if index >= len(nums) {
			return nil
		}
		s := fmt.Sprintf("%v", nums[index])
		if s == "null" {
			return nil
		} else {
			atoi, _ := strconv.Atoi(s)

			r := &TreeNode{Val: atoi}
			r.Left = build(nums, 2*index+1)
			r.Right = build(nums, 2*index+2)
			return r
		}
	}

	return build(nums, 0)
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
