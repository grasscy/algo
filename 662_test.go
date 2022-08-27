package algo

import "testing"

func Test_widthOfBinaryTree(t *testing.T) {
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
	n6 := &TreeNode{
		Val:   6,
		Left:  nil,
		Right: nil,
	}
	n5 := &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	n9 := &TreeNode{
		Val:   9,
		Left:  nil,
		Right: nil,
	}
	n7 := &TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}
	n1.Left = n3
	n1.Right = n2
	n3.Left = n5
	n5.Left = n6
	n2.Right = n9
	n9.Left = n7
	widthOfBinaryTree(n1)
}
