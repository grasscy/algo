package ac

func numComponents(head *ListNode, nums []int) int {
	if head == nil {
		return 0
	}
	s := map[int]bool{}
	for _, v := range nums {
		s[v] = true
	}
	p := head
	ans := 0
	for p != nil {
		if s[p.Val] {
			for p != nil && s[p.Val] {
				p = p.Next
			}
			ans++
		} else {
			p = p.Next
		}
	}
	return ans
}
