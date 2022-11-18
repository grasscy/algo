package pass

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	ans := [][]int{}
	q := []*TreeNode{}
	q = append(q, root)
	h := 1
	for len(q) != 0 {
		n := len(q)
		temp := []int{}
		for i := 0; i < n; i++ {
			t := q[0]
			q = q[1:]
			temp = append(temp, t.Val)
			if t.Left != nil {
				q = append(q, t.Left)
			}
			if t.Right != nil {
				q = append(q, t.Right)
			}
		}
		tt := []int{}
		if h%2 == 0 {
			for i := len(temp) - 1; i >= 0; i-- {
				tt = append(tt, temp[i])
			}
		} else {
			tt = temp
		}
		h++
		ans = append(ans, tt)
	}
	return ans
}
