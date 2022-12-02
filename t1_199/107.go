package t1_199

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	q := []*TreeNode{root}
	ans := [][]int{}
	for len(q) != 0 {
		n := len(q)
		tt := []int{}
		for i := 0; i < n; i++ {
			t := q[0]
			tt = append(tt, t.Val)
			q = q[1:]
			if t.Left != nil {
				q = append(q, t.Left)
			}
			if t.Right != nil {
				q = append(q, t.Right)
			}
		}
		ans = append(ans, tt)
	}
	ret := [][]int{}
	for i := len(ans) - 1; i >= 0; i-- {
		ret = append(ret, ans[i])
	}
	return ret
}
