package t1_199

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	for true {
		if fast == nil || fast.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	if slow == fast {
		fast = head
		for slow != fast {
			slow = slow.Next
			fast = fast.Next
		}
		return fast
	}

	return nil
}
