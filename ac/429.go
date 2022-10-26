package ac

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	q := []*Node{root}
	for len(q) != 0 {
		n := len(q)
		r := []int{}
		for i := 0; i < n; i++ {
			t := q[0]
			q = q[1:]
			r = append(r, t.Val)
			for _, v := range t.Children {
				if v != nil {
					q = append(q, v)
				}
			}
		}
		ans = append(ans, r)
	}
	return ans
}
