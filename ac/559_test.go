package ac

import "testing"

func Test_maxDepth(t *testing.T) {
	n1 := &Node{Val: 1}
	n2 := &Node{Val: 2}
	n3 := &Node{Val: 3}
	n4 := &Node{Val: 4}
	n5 := &Node{Val: 5}
	n6 := &Node{Val: 6}

	n1.Children = []*Node{n3, n2, n4}
	n3.Children = []*Node{n5, n6}

	maxDepth(n1)
}
