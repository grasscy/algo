package pass

import "testing"

func Test_ff(t *testing.T) {
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
	n1.Left = n2
	n2.Left = n3
	n2.Right = n4
	n1.Right = n5
	n5.Right = n6

	flatten(n1)
}
