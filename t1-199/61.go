package pass

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	n := 1
	cur := head
	for cur.Next != nil {
		cur = cur.Next
		n++
	}

	k = k % n
	tail := cur
	tail.Next = head

	cur = head
	for i := 1; i < n-k; i++ {
		cur = cur.Next
	}
	next := cur.Next
	cur.Next = nil
	return next

}
