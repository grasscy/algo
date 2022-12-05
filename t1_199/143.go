package t1_199

func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	ns := []*ListNode{}

	cur := head
	for cur != nil {
		next := cur.Next
		ns = append(ns, cur)
		cur.Next = nil
		cur = next

	}

	var pre *ListNode

	i := 0
	for ; i < len(ns)/2; i++ {
		if pre != nil {
			pre.Next = ns[i]
		}
		if len(ns)-1-i != i {
			ns[i].Next = ns[len(ns)-1-i]
			pre = ns[len(ns)-1-i]
		}
	}
	if len(ns)%2 == 1 && pre != nil {
		pre.Next = ns[i]
	}

	return
}
