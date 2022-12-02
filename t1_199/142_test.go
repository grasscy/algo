package t1_199

import (
	"testing"
)

func Test_detectCycle(t *testing.T) {
	n1 := &ListNode{3, nil}
	n2 := &ListNode{2, nil}
	n3 := &ListNode{0, nil}
	n4 := &ListNode{-1, nil}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n2
	detectCycle(n1)

}
