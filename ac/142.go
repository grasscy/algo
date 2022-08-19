package ac

func detectCycle(head *ListNode) *ListNode {
    vd := map[*ListNode]bool{}
    if head == nil {
        return nil
    }
    for head != nil {
        if vd[head] {
            return head
        } else {
            vd[head] = true
            head = head.Next
        }
    }
    return nil
}
