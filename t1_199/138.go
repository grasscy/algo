package t1_199

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

var m map[*Node]*Node //old->new

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	m = map[*Node]*Node{}

	cur := head

	for cur != nil {
		nn := &Node{cur.Val, nil, nil}
		m[cur] = nn
		cur = cur.Next
	}
	cur = head
	for cur != nil {
		nn := m[cur]
		nn.Next = m[cur.Next]
		nn.Random = m[cur.Random]
		cur = cur.Next
	}
	return m[head]
}
