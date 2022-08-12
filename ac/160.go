package ac

func getIntersectionNode(headA, headB *ListNode) *ListNode {
    m := map[*ListNode]*ListNode{}
    for headA != nil {
        m[headA] = headA
        headA = headA.Next
    }

    for headB != nil {
        v, ok := m[headB]
        if ok {
            return v
        }
        headB = headB.Next
    }
    return nil
}
