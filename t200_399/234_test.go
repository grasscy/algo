package t200_399

import "testing"

func Test_isPalindrome(t *testing.T) {
	n1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	n2 := &ListNode{
		Val:  2,
		Next: nil,
	}
	n1.Next = n2
	isPalindrome(n1)
}
