package t1_199

type Node struct {
	Val       int
	Neighbors []*Node
}

var m map[int]*Node
var mold map[int]*Node

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	m = map[int]*Node{}
	mold = map[int]*Node{}
	d(node)

	for _, v := range mold {
		f(v)
	}

	return m[node.Val]
}

func f(node *Node) {
	nn := m[node.Val]
	nn.Neighbors = []*Node{}
	for _, v := range node.Neighbors {
		nn.Neighbors = append(nn.Neighbors, m[v.Val])
	}
}

func d(node *Node) {
	if node == nil {
		return
	}
	if m[node.Val] != nil {
		return
	}
	mold[node.Val] = node
	nn := &Node{node.Val, nil}
	m[node.Val] = nn
	for _, v := range node.Neighbors {
		if m[v.Val] == nil {
			d(v)
		}
	}
}
