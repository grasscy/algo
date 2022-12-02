package t1_199

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	q := []*Node{root}
	for len(q) != 0 {
		n := len(q)
		var pre *Node
		for i := 0; i < n; i++ {
			t := q[0]
			q = q[1:]
			if pre != nil {
				pre.Next = t
			}
			pre = t
			if t.Left != nil {
				q = append(q, t.Left)
			}
			if t.Right != nil {
				q = append(q, t.Right)
			}
		}
	}
	return root
}
