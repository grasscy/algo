package algo

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    ret := &ListNode{}
    cur := ret

    c1 := l1
    c2 := l2

    v1 := c1.Val
    v2 := c2.Val
    r := v1 + v2
    cur.Val = r % 10
    yu := r / 10
    for c1.Next != nil || c2.Next != nil || yu != 0 {
        v1 = 0
        v2 = 0
        if c1.Next != nil {
            c1 = c1.Next
            v1 = c1.Val
        }
        if c2.Next != nil {
            c2 = c2.Next
            v2 = c2.Val
        }
        r := v1 + v2 + yu
        yu = r / 10
        n := ListNode{
            Val:  r % 10,
            Next: nil,
        }
        cur.Next = &n
        cur = cur.Next
    }
    return ret
}
