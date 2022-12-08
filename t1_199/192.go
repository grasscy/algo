package t1_199

func rightSideView(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		ans = append(ans, q[len(q)-1].Val)
		n := len(q)
		for i := 0; i < n; i++ {
			t := q[0]
			q = q[1:]
			if t.Left != nil {
				q = append(q, t.Left)
			}
			if t.Right != nil {
				q = append(q, t.Right)
			}
		}
	}
	return ans
}
