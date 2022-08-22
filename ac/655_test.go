package ac

import (
	"testing"
)

func Test_printTree(t *testing.T) {
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

	n1.Left = n2
	printTree(n1)
	n2.Right = n4
	n1.Right = n3
	printTree(n1)


}
