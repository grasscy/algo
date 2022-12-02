package t1_199

import (
	"testing"
)

func Test_levelOrder(t *testing.T) {
	n3 := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	n9 := &TreeNode{
		Val:   9,
		Left:  nil,
		Right: nil,
	}
	n20 := &TreeNode{
		Val:   20,
		Left:  nil,
		Right: nil,
	}
	n15 := &TreeNode{
		Val:   15,
		Left:  nil,
		Right: nil,
	}
	n7 := &TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}

	n3.Left = n9
	n3.Right = n20
	n20.Left = n15
	n20.Right = n7
	levelOrder(n3)

}
