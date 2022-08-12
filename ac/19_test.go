package ac

import (
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	n1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	//n2 := &ListNode{
	//	Val:  2,
	//	Next: nil,
	//}
	//n3 := &ListNode{
	//	Val:  3,
	//	Next: nil,
	//}
	//n4 := &ListNode{
	//	Val:  4,
	//	Next: nil,
	//}
	//n5 := &ListNode{
	//	Val:  5,
	//	Next: nil,
	//}

	//n1.Next = n2
	//n2.Next = n3
	//n3.Next = n4
	//n4.Next = n5
	removeNthFromEnd(n1,1)
}
