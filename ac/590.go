package ac

func postorder(root *Node) []int {
	if root == nil {
		return []int{}
	}
	ans := []int{}
	for _, v := range root.Children {
		ans = append(ans, postorder(v)...)
	}
	ans = append(ans, root.Val)
	return ans
}

type Node struct {
	Val      int
	Children []*Node
}
