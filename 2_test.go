package algo

import (
    "testing"
)

func Test_addTwoNumbers(t *testing.T) {
    n1 := &ListNode{
        Val:  2,
        Next: nil,
    }
    n2 := &ListNode{
        Val:  4,
        Next: nil,
    }
    n3 := &ListNode{
        Val:  3,
        Next: nil,
    }
    n4 := &ListNode{
        Val:  5,
        Next: nil,
    }
    n5 := &ListNode{
        Val:  6,
        Next: nil,
    }
    n6 := &ListNode{
        Val:  4,
        Next: nil,
    }

    n1.Next = n2
    n2.Next = n3
    n4.Next = n5
    n5.Next = n6

    addTwoNumbers(n1, n4)

}
