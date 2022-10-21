package ac

func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	q := []*TreeNode{}
	q = append(q, root)
	ans := []int{}
	for len(q) != 0 {
		n := len(q)
		tmpm := q[0].Val
		for i := 0; i < n; i++ {
			t := q[0]
			q = q[1:]
			if t.Val > tmpm {
				tmpm = t.Val
			}
			if t.Left != nil {
				q = append(q, t.Left)
			}
			if t.Right != nil {
				q = append(q, t.Right)
			}
		}
		ans = append(ans, tmpm)
	}
	return ans
}
