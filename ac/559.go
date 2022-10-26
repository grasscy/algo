package ac

var ans int

func maxDepth(root *Node) int {
	ans = 0
	return d(root)
}

func d(root *Node) int {
	if root == nil {
		return 0
	}
	md := 1
	for _, v := range root.Children {
		md = max(md, d(v)+1)
	}
	ans = max(md, ans)
	return md
}
