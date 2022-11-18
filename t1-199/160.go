package pass

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pa := headA
	pb := headB
	for pa != pb {
		if pa != nil {
			pa = pa.Next
		} else {
			pa = headB
		}
		if pb != nil {
			pb = pb.Next
		} else {
			pb = headA
		}
	}
	return pb

}
