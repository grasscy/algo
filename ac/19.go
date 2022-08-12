package ac

//opt
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    l := 0
    cur := head
    for cur != nil {
        l++
        cur = cur.Next
    }
    dummy := &ListNode{0, head}
    cur = dummy
    for i := 0; i < l-n; i++ {
        cur = cur.Next
    }
    cur.Next = cur.Next.Next
    return dummy.Next
}
