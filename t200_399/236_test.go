package t200_399

import (
	"fmt"
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	n0 := &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}
	n1 := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	n2 := &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	n3 := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	n4 := &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	n5 := &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	n6 := &TreeNode{
		Val:   6,
		Left:  nil,
		Right: nil,
	}
	n7 := &TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}
	n8 := &TreeNode{
		Val:   8,
		Left:  nil,
		Right: nil,
	}
	n3.Left = n5
	n5.Left = n6
	n5.Right = n2
	n2.Left = n7
	n2.Right = n4
	n3.Right = n1
	n1.Left = n0
	n1.Right = n8
	ancestor := lowestCommonAncestor(n3, n5, n4)
	fmt.Println(ancestor)
}
